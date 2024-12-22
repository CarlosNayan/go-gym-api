package main

import (
	"fmt"
	"os"
	"time"

	"api-gym-on-go/models"
	"api-gym-on-go/src/config/env"
	"api-gym-on-go/src/modules/auth"
	"api-gym-on-go/src/modules/checkins"
	"api-gym-on-go/src/modules/gyms"
	"api-gym-on-go/src/modules/users"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	env.LoadEnv()

	// Startup services
	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	})

	app.Use(logger.New())
	db := models.SetupDatabase(env.DatabaseURL)

	// Register modules
	auth.Register(app, db)
	users.Register(app, db)
	gyms.Register(app, db)
	checkins.Register(app, db)

	// Start server
	port := 3000
	if env.Port != 0 {
		port = env.Port
	}

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
		os.Exit(1)
	}
}
