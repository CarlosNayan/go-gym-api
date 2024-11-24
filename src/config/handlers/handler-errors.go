package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HTTPError interface {
	error
	StatusCode() int
}

func HandleHTTPError(c *fiber.Ctx, err error) error {
	if httpErr, ok := err.(HTTPError); ok {
		return c.Status(httpErr.StatusCode()).JSON(fiber.Map{
			"error": httpErr.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Internal Server Error",
	})
}
