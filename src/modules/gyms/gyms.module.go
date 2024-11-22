package gyms

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/middleware"
	"api-gym-on-go/src/modules/gyms/repository"
	"api-gym-on-go/src/modules/gyms/services"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app *fiber.App, db *gorm.DB) {
	gymsRepo := repository.NewGymsRepository(db)
	gymCreateService := services.NewGymsCreateService(gymsRepo)
	gymNearbyService := services.NewGymsNearbyService(gymsRepo)
	gymSearchService := services.NewGymsSearchService(gymsRepo)

	app.Post("/gyms/create", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		var gym models.Gym

		if err := c.BodyParser(&gym); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		err := gymCreateService.CreateGym(&gym)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(gym)
	})

	app.Get("/gyms/nearby", func(c *fiber.Ctx) error {
		latitudeStr := c.Query("latitude")
		latitude, err := strconv.ParseFloat(latitudeStr, 64)
		if err != nil || math.Abs(latitude) > 90 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid latitude"})
		}

		longitudeStr := c.Query("longitude")
		longitude, err := strconv.ParseFloat(longitudeStr, 64)
		if err != nil || math.Abs(longitude) > 180 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid longitude"})
		}

		gyms, err := gymNearbyService.GetGymsNearby(latitude, longitude)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(gyms)
	})

	app.Get("/gyms/search", func(c *fiber.Ctx) error {
		query := c.Query("query")

		gyms, err := gymSearchService.SearchGyms(query)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(gyms)
	})
}
