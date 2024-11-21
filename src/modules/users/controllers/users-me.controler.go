package controllers

import (
	"api-gym-on-go/src/modules/users/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (ctrl *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.UserService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}
