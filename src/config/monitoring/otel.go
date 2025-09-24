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
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/encoding/gzip"
)

type Providers struct {
	TracerProvider *sdktrace.TracerProvider
	MeterProvider  *metric.MeterProvider
}

func InitOTEL(ctx context.Context) (*Providers, error) {
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("OTEL_EXPORTER_OTLP_ENDPOINT not set")
	}

	exportInterval := 10 * time.Second
	exportTimeout := 5 * time.Second

	traceExp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithCompressor(gzip.Name),
	)
	if err != nil {
		return nil, err
	}

	metricExp, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(endpoint),
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithCompressor(gzip.Name),
	)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("api-gym-on-go"),
			semconv.ServiceVersion("1.0.0"),
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
		os.Exit(0)
	}()

	return &Providers{TracerProvider: tp, MeterProvider: mp}, nil
}

func InitDB(dsn string) *sql.DB {
	driverName, err := otelsql.Register(
		"postgres",
		otelsql.AllowRoot(),
		otelsql.TraceQueryWithoutArgs(),
		otelsql.TraceRowsClose(),
		otelsql.TraceRowsAffected(),
		otelsql.WithDatabaseName("go-gym-api"),
		otelsql.WithSystem(semconv.DBSystemNamePostgreSQL),
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	if err := otelsql.RecordStats(db); err != nil {
		log.Fatalf("Warning: could not record db stats: %v", err)
	}

	return db
}

func FiberOtelTracingMiddleware(serviceName string) fiber.Handler {
	tracer := otel.Tracer(serviceName)

	return func(c *fiber.Ctx) error {
		methodRaw := c.Method()
		method := string([]byte(methodRaw))

		routeRaw := ""
		if c.Route() != nil && c.Route().Path != "" {
			routeRaw = c.Route().Path
		} else {
			routeRaw = c.Path()
		}
		routePath := string([]byte(routeRaw))

		originalURLRaw := c.OriginalURL()
		originalURL := string([]byte(originalURLRaw))

		ctx, span := tracer.Start(c.Context(),
			fmt.Sprintf("%s %s", method, routePath),
			trace.WithAttributes(
				attribute.String("http.method", method),
				attribute.String("http.route", routePath),
				attribute.String("http.url", originalURL),
			),
		)
		defer span.End()

		c.SetUserContext(ctx)

		err := c.Next()

		if err != nil {
			span.RecordError(err)
		}
		span.SetAttributes(attribute.Int("http.status_code", c.Response().StatusCode()))

		return err
	}
}

func FiberOtelMetricsMiddleware(serviceName string) fiber.Handler {
	meter := otel.Meter(serviceName)

	requestDuration, err := meter.Float64Histogram(
		"http.server.metrics",
		otelMetric.WithUnit("s"),
	)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start).Seconds()

		ctx := c.UserContext()
		if ctx == nil {
			ctx = context.Background()
		}

		methodRaw := c.Method()
		method := string([]byte(methodRaw))

		routeRaw := ""
		if c.Route() != nil && c.Route().Path != "" {
			routeRaw = c.Route().Path
		} else {
			routeRaw = c.Path()
		}
		routePath := string([]byte(routeRaw))

		statusRaw := c.Response().StatusCode()

		attrs := []attribute.KeyValue{
			attribute.String("http.method", method),
			attribute.String("http.route", routePath),
			attribute.Int("http.status_code", statusRaw),
		}

		requestDuration.Record(ctx, duration, otelMetric.WithAttributes(attrs...))

		return err
	}
}
