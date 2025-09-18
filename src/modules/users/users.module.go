package users

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/handlers"
	"api-gym-on-go/src/config/middleware"
	"api-gym-on-go/src/config/validate"
	"api-gym-on-go/src/modules/users/repository"
	users_schemas "api-gym-on-go/src/modules/users/schemas"
	"api-gym-on-go/src/modules/users/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, db *sql.DB) {

	app.Get("/users/me", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_user := c.Locals("sub").(string)

		repo := repository.NewUserRepository(db)
		service := services.NewUsersMeService(repo)

		user, err := service.GetUserByID(id_user)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(user)
	})

	app.Post("/users/create", func(ctx *fiber.Ctx) error {
		body, err := validate.ParseBody[users_schemas.UserCreateBody](ctx)
		if err != nil {
			return handlers.HandleHTTPError(ctx, &errors.CustomError{Message: err.Error(), Code: 400})
		}

		repo := repository.NewUserRepository(db)
		service := services.NewUsersCreateService(repo)

		createdUser, err := service.CreateUser(body)
		if err != nil {
			return handlers.HandleHTTPError(ctx, err)
		}

		return ctx.Status(201).JSON(createdUser)
	})
}
