package main

import (
	"fmt"
	"os"

	"api-gym-on-go/models"
	"api-gym-on-go/src/config/env"
	"api-gym-on-go/src/modules/auth"
	"api-gym-on-go/src/modules/gyms"
	"api-gym-on-go/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	envConfig, err := env.LoadEnv()
	if err != nil {
		fmt.Printf("Erro ao carregar variáveis de ambiente: %v\n", err)
		os.Exit(1)
	}

	// Startup services
	app := fiber.New()
	db := models.SetupDatabase(envConfig.DatabaseURL)

	// Register modules
	users.Register(app, db)
	auth.Register(app, db)
	gyms.Register(app, db)

	// Start server
	port := 3000
	if envConfig.Port != 0 {
		port = envConfig.Port
	}

	err = app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
		os.Exit(1)
	}
}
