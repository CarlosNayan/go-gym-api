package auth

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/handlers"
	"api-gym-on-go/src/config/validate"
	"api-gym-on-go/src/modules/auth/repository"
	auth_schemas "api-gym-on-go/src/modules/auth/schemas"
	"api-gym-on-go/src/modules/auth/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(app *fiber.App, db *sql.DB) {
	app.Post("/auth", func(ctx *fiber.Ctx) error {
		authBody, err := validate.ParseBody[auth_schemas.AuthRequestBody](ctx)
		if err != nil {
			return handlers.HandleHTTPError(ctx, &errors.CustomError{Message: err.Error(), Code: 400})
		}

		repo := repository.NewAuthRepository(db)
		service := services.NewAuthService(repo)

		user, err := service.Execute(authBody.Email, authBody.Password)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.JSON(user)
	})
}
