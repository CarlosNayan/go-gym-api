package monitoring

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"go.nhat.io/otelsql"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	otelMetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/encoding/gzip"
)

type Providers struct {
	TracerProvider *sdktrace.TracerProvider
	MeterProvider  *metric.MeterProvider
}

// MetricJob representa uma métrica a ser registrada
type MetricJob struct {
	Ctx        context.Context // <- context.Context
	Method     string
	Route      string
	StatusCode int
	Duration   float64
}

// TraceJob representa um span a ser finalizado
type TraceJob struct {
	Ctx        context.Context
	Span       trace.Span
	StatusCode int
	Err        error
}

// Fila de métricas
var metricQueue chan MetricJob

// Fila de traces
var traceQueue chan TraceJob

// InitOTEL inicializa os providers de métricas e traces
func InitOTEL(ctx context.Context) (*Providers, error) {
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("OTEL_EXPORTER_OTLP_ENDPOINT not set")
	}

	// Intervalo e timeout configuráveis
	exportInterval := 10 * time.Second
	exportTimeout := 5 * time.Second

	// Trace exporter (gRPC + GZIP)
	traceExp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithInsecure(), // trocar para TLS em produção
		otlptracegrpc.WithCompressor(gzip.Name),
	)
	if err != nil {
		return nil, err
	}

	// Metric exporter (gRPC + GZIP)
	metricExp, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(endpoint),
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithCompressor(gzip.Name),
	)
	if err != nil {
		return nil, err
	}

	// Resource metadata
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("api-gym-on-go"),
			semconv.ServiceVersion("1.0.0"),
			semconv.DeploymentEnvironment(os.Getenv("DEPLOY_ENV")),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExp),
		sdktrace.WithResource(res),
	)

	mp := metric.NewMeterProvider(
		metric.WithReader(
			metric.NewPeriodicReader(
				metricExp,
				metric.WithInterval(exportInterval),
				metric.WithTimeout(exportTimeout),
			),
		),
		metric.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(mp)

	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		log.Println("Shutting down OpenTelemetry providers...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
		if err := mp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down meter provider: %v", err)
		}
		os.Exit(0) // encerra a aplicação
	}()

	return &Providers{TracerProvider: tp, MeterProvider: mp}, nil
}

// InitDB registra o driver do Postgres com métricas
func InitDB(dsn string) *sql.DB {
	driverName, err := otelsql.Register(
		"postgres",
		otelsql.WithSystem(semconv.DBSystemPostgreSQL),
	)
	if err != nil {
		log.Fatalf("Warning: could not register db driver: %v", err)
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatalf("Warning: could not open db connection: %v", err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Testa conexão imediatamente
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	// Registrar métricas de stats do DB
	if err := otelsql.RecordStats(db); err != nil {
		log.Fatalf("Warning: could not record db stats: %v", err)
	}

	return db
}

func FiberOtelMiddleware(serviceName string) fiber.Handler {
	tracer := otel.Tracer(serviceName)

	return func(c *fiber.Ctx) error {
		ctx, span := tracer.Start(c.Context(),
			c.Method()+" "+c.Path(),
			trace.WithAttributes(
				attribute.String("http.method", c.Method()),
				attribute.String("http.route", c.Route().Path),
				attribute.String("http.url", c.OriginalURL()),
			),
		)
		defer span.End()

		// injeta o ctx no Fiber
		c.SetUserContext(ctx)

		err := c.Next()
		if err != nil {
			span.RecordError(err)
		}
		span.SetAttributes(attribute.Int("http.status_code", c.Response().StatusCode()))
		return err
	}
}

// FiberMetricsMiddleware cria um middleware que registra métricas HTTP.
func FiberMetricsMiddleware() fiber.Handler {
	meter := otel.Meter("api-gym-on-go")

	requestDuration, err := meter.Float64Histogram(
		"http.server.duration",
		otelMetric.WithUnit("s"),
	)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start).Seconds() // segundos

		ctx := c.UserContext()
		attrs := []attribute.KeyValue{
			attribute.String("http.method", c.Method()),
			attribute.String("http.route", c.Route().Path),
			attribute.Int("http.status_code", c.Response().StatusCode()),
		}

		requestDuration.Record(ctx, duration, otelMetric.WithAttributes(attrs...))

		return err
	}
}
