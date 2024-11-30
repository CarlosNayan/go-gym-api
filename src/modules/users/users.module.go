package users

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/handlers"
	"api-gym-on-go/src/config/middleware"
	"api-gym-on-go/src/modules/users/repository"
	"api-gym-on-go/src/modules/users/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, db *sql.DB) {
	userRepo := repository.NewUserRepository(db)
	usersMeService := services.NewUsersMeService(userRepo)
	usersCreateService := services.NewUsersCreateService(userRepo)

	app.Get("/users/me", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_user := c.Locals("sub").(string)

		user, err := usersMeService.GetUserByID(id_user)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(user)
	})

	app.Post("/users/create", func(c *fiber.Ctx) error {
		var user models.User

		if err := c.BodyParser(&user); err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid request body",
				Code:    fiber.StatusBadRequest,
			})
		}

		createdUser, err := usersCreateService.CreateUser(&user)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.Status(201).JSON(createdUser)
	})
}
