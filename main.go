package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"api-gym-on-go/src/config/env"
	"api-gym-on-go/src/config/monitoring"
	"api-gym-on-go/src/config/utils"
	"api-gym-on-go/src/modules/auth"
	"api-gym-on-go/src/modules/checkins"
	"api-gym-on-go/src/modules/gyms"
	"api-gym-on-go/src/modules/users"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	env.LoadEnv()

	// OTEL context
	ctx := context.Background()

	// OTEL startup
	provs, err := monitoring.InitOTEL(ctx)
	if err != nil {
		log.Fatalf("failed to init OTEL: %v", err)
	}
	// Shutdown OTEL providers
	defer func() {
		_ = provs.TracerProvider.Shutdown(ctx)
		_ = provs.MeterProvider.Shutdown(ctx)
	}()

	// Start DB instrumented
	db := monitoring.InitDB(env.DATABASE_URL)

	// Startup services
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(monitoring.FiberOtelMiddleware("api-gym-on-go"))
	app.Use(monitoring.FiberMetricsMiddleware())

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
			// Ou valide origens específicas
			// return origin == "http://localhost:3000" || origin == "http://192.168.x.x:19000"
		},
		AllowMethods: strings.Join([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}, ","),
		AllowHeaders: strings.Join([]string{
			"Content-Type",
			"Content-Length",
			"Accept",
			"Content-Range",
			"User-Offset",
			"User-Platform",
			"User-Device",
			"User-Os-Name",
			"User-Os-Version",
			"Authorization",
			"Sec-WebSocket-Protocol",
			"Sec-Websocket-Key",
			"Sec-Websocket-Extensions",
			"Sec-Websocket-Version",
			"CF-Connecting-IP",
			"CF-IPCountry",
			"CF-Ray",
			"CF-Visitor",
			"True-Client-IP",
		}, ","),
		AllowCredentials: true,
		MaxAge:           int(120 * time.Hour.Seconds()),
	}))

	app.Use(logger.New(logger.Config{
		Format:        "${time} | ${latency} | ${status} | ${method} | ${path}\n",
		TimeFormat:    "2006-01-02 15:04:05",
		TimeZone:      "UTC",
		DisableColors: true,
	}))

	// Register modules
	auth.Register(app, db)
	users.Register(app, db)
	gyms.Register(app, db)
	checkins.Register(app, db)

	utils.RouteLogger(app, env.PORT)

	err = app.Listen(fmt.Sprintf(":%d", env.PORT))
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
		os.Exit(1)
	}
}
