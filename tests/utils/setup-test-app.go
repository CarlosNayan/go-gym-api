package utils

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/auth"
	"api-gym-on-go/src/modules/checkins"
	"api-gym-on-go/src/modules/gyms"
	"api-gym-on-go/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func SetupTestApp() *fiber.App {
	app := fiber.New()
	db := models.SetupDatabase("postgresql://root:admin@127.0.0.1:5432/api_solid?sslmode=disable")


	users.Register(app, db)
	auth.Register(app, db)
	gyms.Register(app, db)
	checkins.Register(app, db)
	return app
}
