package users

import (
	"api-gym-on-go/src/modules/users/controllers"
	"api-gym-on-go/src/modules/users/repository"
	"api-gym-on-go/src/modules/users/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app *fiber.App, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	app.Get("/users/:id", userController.GetUser)
}
