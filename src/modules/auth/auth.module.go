package auth

import (
	"api-gym-on-go/src/modules/auth/repository"
	"api-gym-on-go/src/modules/auth/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(app *fiber.App, db *sql.DB, jwtSecret *string) {
	jwtKey := []byte(*jwtSecret)

	authService := repository.NewAuthRepository(db)
	userService := services.NewAuthService(authService, jwtKey)

	app.Post("/auth", func(c *fiber.Ctx) error {
		var authRequest AuthRequest

		if err := c.BodyParser(&authRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		email := authRequest.Email
		password := authRequest.Password

		user, err := userService.Auth(email, password)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(user)
	})
}
