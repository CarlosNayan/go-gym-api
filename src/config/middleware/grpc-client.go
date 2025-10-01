package middleware

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api-gym-on-go/src/config/env"
	monitoringpb "api-gym-on-go/src/config/helpers/monitoringpb"
)

var grpcClient monitoringpb.MonitoringServiceClient

func InitGrpcClient() {

	conn, err := grpc.NewClient(env.GRPC_METRICS_EXPORTER_ENDPOINT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Erro ao conectar no servidor gRPC: %v", err)
	}
	grpcClient = monitoringpb.NewMonitoringServiceClient(conn)
}

func GrpcMetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start).Seconds() * 1000

		methodRaw := c.Method()
		method := string([]byte(methodRaw))

		routeRaw := ""
		if c.Route() != nil && c.Route().Path != "" {
			routeRaw = c.Route().Path
		} else {
			routeRaw = c.Path()
		}
		routePath := string([]byte(routeRaw))

		originalUrlRaw := c.OriginalURL()
		originalUrl := string([]byte(originalUrlRaw))

		statusRaw := c.Response().StatusCode()
		status := strconv.Itoa(statusRaw)

		metric := &monitoringpb.HttpMetric{
			HttpMethod:     method,
			HttpRoute:      routePath,
			OriginalUrl:    originalUrl,
			HttpStatusCode: status,
			DurationMs:     duration,
			Instance:       "api-instance-1",
			CreatedAtUnix:  time.Now().Unix(),
		}

		go func(m *monitoringpb.HttpMetric) {
			_, err := grpcClient.SendHttpMetric(context.Background(), m)
			if err != nil {
				log.Printf("Erro ao enviar métrica gRPC: %v", err)
			}
		}(metric)

		return err
	}
}
