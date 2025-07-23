package gyms

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/handlers"
	"api-gym-on-go/src/config/middleware"
	"api-gym-on-go/src/models"
	"api-gym-on-go/src/modules/gyms/repository"
	"api-gym-on-go/src/modules/gyms/services"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	gymsRepo := repository.NewGymsRepository()
	gymCreateService := services.NewGymsCreateService(gymsRepo)
	gymNearbyService := services.NewGymsNearbyService(gymsRepo)
	gymSearchService := services.NewGymsSearchService(gymsRepo)

	app.Post("/gyms/create", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		var gym models.Gym

		if err := c.BodyParser(&gym); err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		err := gymCreateService.CreateGym(&gym)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(gym)
	})

	app.Get("/gyms/nearby", func(c *fiber.Ctx) error {
		latitudeStr := c.Query("latitude")
		latitude, err := strconv.ParseFloat(latitudeStr, 64)
		if err != nil || math.Abs(latitude) > 90 {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid latitude",
				Code:    fiber.StatusBadRequest,
			})
		}

		longitudeStr := c.Query("longitude")
		longitude, err := strconv.ParseFloat(longitudeStr, 64)
		if err != nil || math.Abs(longitude) > 180 {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid longitude",
				Code:    fiber.StatusBadRequest,
			})
		}

		gyms, err := gymNearbyService.GetGymsNearby(latitude, longitude)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(gyms)
	})

	app.Get("/gyms/search", func(c *fiber.Ctx) error {
		query := c.Query("query")

		gyms, err := gymSearchService.SearchGyms(query)
		if err != nil {
			fmt.Println(err)
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(gyms)
	})
}
