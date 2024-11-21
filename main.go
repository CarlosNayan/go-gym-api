package main

import (
	"fmt"
	"os"

	"api-gym-on-go/schema"
	"api-gym-on-go/src/config/env"
	"api-gym-on-go/src/modules/auth"
	"api-gym-on-go/src/modules/users"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

func main() {

	envConfig, err := env.LoadEnv()
	if err != nil {
		fmt.Printf("Erro ao carregar variáveis de ambiente: %v\n", err)
		os.Exit(1)
	}

	app := fiber.New()

	db := schema.SetupDatabase(envConfig.DatabaseURL)

	users.Register(app, db)
	auth.Register(app, db)

	port := 3000
	if envConfig.Port != 0 {
		port = envConfig.Port
	}

	err = app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
