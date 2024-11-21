package users

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
}
