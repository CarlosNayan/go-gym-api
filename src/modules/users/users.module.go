package users

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/middleware"
	"api-gym-on-go/src/modules/users/repository"
	"api-gym-on-go/src/modules/users/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app *fiber.App, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	usersMeService := services.NewUserMeService(userRepo)
	usersCreateService := services.NewUsersCreateService(userRepo)

	app.Get("/users/me", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_user := c.Locals("sub").(string)

		user, err := usersMeService.GetUserByID(id_user)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(user)
	})

	app.Post("/users/create", func(c *fiber.Ctx) error {
		var user models.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		err := usersCreateService.CreateUser(&user)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(user)
	})
}
