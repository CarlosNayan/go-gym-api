// src/config/monitoring/monitoring.go
package monitoring

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"go.nhat.io/otelsql"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
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

	// Trace exporter
	traceExp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithInsecure(), // trocar para TLS em produção
	)
	if err != nil {
		return nil, err
	}

	// Metric exporter
	metricExp, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(endpoint),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("api-gym-on-go"),
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
			metric.NewPeriodicReader(metricExp, metric.WithInterval(10*time.Second)),
		),
		metric.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(mp)

	return &Providers{TracerProvider: tp, MeterProvider: mp}, nil
}

func InitDB(dsn string) *sql.DB {
	driverName, err := otelsql.Register("postgres",
		otelsql.WithSystem(semconv.DBSystemPostgreSQL),
	)
	if err != nil {
		log.Fatalf("Warning: could not register db driver: %v", err)
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatalf("Warning: could not open db connection: %v", err)
	}
	// Registrar métricas de stats do DB, se quiser
	if err := otelsql.RecordStats(db); err != nil {
		log.Fatalf("Warning: could not record db stats: %v", err)
	}

	return db
}

// // MiddlewareFiber retorna o handler de middleware do Fiber já configurado
// func MiddlewareFiber() fiber.Handler {
// 	return otelfiber.Middleware("api-gym-on-go")
// }
