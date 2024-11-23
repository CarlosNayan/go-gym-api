package checkins

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/middleware"
	"api-gym-on-go/src/modules/checkins/repository"
	"api-gym-on-go/src/modules/checkins/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app *fiber.App, db *gorm.DB) {
	checkinRepo := repository.NewCheckinRepository(db)
	checkinValidateService := services.NewCheckinValidateService(checkinRepo)
	checkinCountHistoryService := services.NewCheckinCountHistory(checkinRepo)
	checkinListHistoryService := services.NewCheckinListHistory(checkinRepo)
	checkinCreateService := services.NewCheckinCreateService(checkinRepo)

	app.Post("/checkin/create", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		var checkin models.Checkin
		if err := c.BodyParser(&checkin); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		err := checkinCreateService.CreateCheckin(&checkin)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(checkin)
	})

	app.Get("/checkin/validate/:id_checkin", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_checkin := c.Params("id_checkin")

		checkin, err := checkinValidateService.ValidateCheckin(id_checkin)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(checkin)
	})
	
	app.Get("/checkin/history/:id_user", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_user := c.Params("id_user")

		count, err := checkinCountHistoryService.CountCheckinHistory(id_user)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(count)
	})

	app.Get("/checkin/history/:id_user", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_user := c.Params("id_user")
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page number"})
		}

		checkins, err := checkinListHistoryService.ListCheckinHistory(id_user, page)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(checkins)
	})
}
