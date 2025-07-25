package utils

import (
	"api-gym-on-go/src/config/database"
	"api-gym-on-go/src/modules/auth"
	"api-gym-on-go/src/modules/checkins"
	"api-gym-on-go/src/modules/gyms"
	"api-gym-on-go/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

// Define os módulos disponíveis como constantes do tipo string
type Module string

const (
	Users    Module = "users"
	Auth     Module = "auth"
	Gyms     Module = "gyms"
	Checkins Module = "checkins"
)

func SetupTestApp(module Module) *fiber.App {
	// Cria uma nova instância do Fiber
	app := fiber.New()

	// Configura o banco de dados
	database.SetupDatabase("postgresql://root:admin@127.0.0.1:5432/public?sslmode=disable")

	// Registra apenas o módulo especificado
	switch module {
	case Users:
		users.Register(app)
	case Auth:
		auth.Register(app)
	case Gyms:
		gyms.Register(app)
	case Checkins:
		checkins.Register(app)
	default:
		panic("Módulo inválido: " + string(module))
	}

	return app
}
