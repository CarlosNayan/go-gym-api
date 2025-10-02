package middleware

import (
	"context"
	"encoding/json"
	"fmt"
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

		// Request tracing and metrics
		reqHeader := make(map[string]string)
		var reqHeaderStr string

		for k, v := range c.Request().Header.All() {
			reqHeader[string(k)] = string(v)
		}

		jsonBytes, err := json.Marshal(reqHeader)
		if err != nil {
			reqHeaderStr = fmt.Sprintf("Erro ao serializar Request Headers para JSON: %v", err)
		} else {
			reqHeaderStr = string(jsonBytes)
		}

		reqBody := string([]byte(c.Request().Body()))

		duration := time.Since(start).Seconds() * 1000

		err = c.Next()

		// Response tracing and metrics
		var errorMessage string
		if err != nil {
			errorMessage = err.Error()
		}

		resHeader := make(map[string]string)
		var resHeaderStr string

		for k, v := range c.Request().Header.All() {
			resHeader[string(k)] = string(v)
		}

		jsonBytes, err = json.Marshal(resHeader)
		if err != nil {
			resHeaderStr = fmt.Sprintf("Erro ao serializar Request Headers para JSON: %v", err)
		} else {
			resHeaderStr = string(jsonBytes)
		}

		resBody := string([]byte(c.Response().Body()))

		routeRaw := ""
		if c.Route() != nil && c.Route().Path != "" {
			routeRaw = c.Route().Path
		} else {
			routeRaw = c.Path()
		}

		method := string([]byte(c.Method()))
		routePath := string([]byte(routeRaw))
		originalUrl := string([]byte(c.OriginalURL()))
		status := strconv.Itoa(c.Response().StatusCode())

		// Campos de Métrica
		metric := &monitoringpb.HttpMetric{
			// Campos de Métrica Original
			HttpMethod:     method,
			HttpRoute:      routePath,
			OriginalUrl:    originalUrl,
			HttpStatusCode: status,
			DurationMs:     duration,
			Instance:       "api-instance",

			// Campos de Tracing (Span)
			StartTimeUnix:   time.Now().Unix(),
			RequestHeaders:  reqHeaderStr,
			RequestBody:     sanitizeBody(reqBody),
			ResponseHeaders: resHeaderStr,
			ResponseBody:    sanitizeBody(resBody),
			ErrorMessage:    errorMessage,
		}

		go func(m *monitoringpb.HttpMetric) {
			_, grpcErr := grpcClient.SendHttpMetric(context.Background(), m)
			if grpcErr != nil {
				log.Printf("Erro ao enviar métrica/trace gRPC: %v", grpcErr)
			}
		}(metric)

		return err
	}
}

func sanitizeBody(body string) string {
	if body == "" {
		return body
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return body
	}

	sensitiveKeys := []string{"password", "new_password", "confirm_password", "old_password", "token", "access_token", "refresh_token"}

	for _, key := range sensitiveKeys {
		if _, ok := data[key]; ok {
			data[key] = ""
		}
	}

	safeBody, err := json.Marshal(data)
	if err != nil {
		return body
	}

	return string(safeBody)
}
