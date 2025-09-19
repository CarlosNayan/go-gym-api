package gyms

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/handlers"
	"api-gym-on-go/src/config/middleware"
	"api-gym-on-go/src/config/validate"
	"api-gym-on-go/src/modules/gyms/repository"
	gyms_schemas "api-gym-on-go/src/modules/gyms/schemas"
	"api-gym-on-go/src/modules/gyms/services"
	"database/sql"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, db *sql.DB) {
	app.Post("/gyms/create", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		body, err := validate.ParseBody[gyms_schemas.GymCreateBody](c)
		if err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{Message: err.Error(), Code: 400})
		}

		repo := repository.NewGymsRepository(db)
		service := services.NewGymsCreateService(repo)

		gym, err := service.CreateGym(body)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(gym)
	})

	app.Get("/gyms/nearby", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		latitudeStr := c.Query("latitude")
		latitude, err := strconv.ParseFloat(latitudeStr, 64)
		if err != nil || math.Abs(latitude) > 90 {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid latitude",
				Code:    400,
			})
		}

		longitudeStr := c.Query("longitude")
		longitude, err := strconv.ParseFloat(longitudeStr, 64)
		if err != nil || math.Abs(longitude) > 180 {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid longitude",
				Code:    400,
			})
		}

		repo := repository.NewGymsRepository(db)
		service := services.NewGymsNearbyService(repo)

		gyms, err := service.GetGymsNearby(latitude, longitude)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(gyms)
	})

	app.Get("/gyms/search", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		query, err := validate.ParseQueryParams[gyms_schemas.GymsSearchQuery](c)
		if err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{Message: err.Error(), Code: 400})
		}

		repo := repository.NewGymsRepository(db)
		service := services.NewGymsSearchService(repo)

		gyms, err := service.SearchGyms(query.Query)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(gyms)
	})
}
