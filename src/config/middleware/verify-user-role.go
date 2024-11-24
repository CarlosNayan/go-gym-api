package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func VerifyUserRole(roleToVerify string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)

		if role != roleToVerify {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		return c.Next()
	}
}
