package checkins

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/handlers"
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
			return handlers.HandleHTTPError(c, err)
		}

		IDUser := c.Locals("sub").(string)

		err := checkinCreateService.CreateCheckin(IDUser, &checkin)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		checkinResponse := models.Checkin{
			ID:          checkin.ID,
			CreatedAt:   checkin.CreatedAt,
			ValidatedAt: checkin.ValidatedAt,
			IDUser:      checkin.IDUser,
			IDGym:       checkin.IDGym,
		}

		return c.JSON(checkinResponse)
	})

	app.Put("/checkin/validate/:id_checkin",
		middleware.ValidateJWT,
		middleware.VerifyUserRole("ADMIN"),
		func(c *fiber.Ctx) error {
			id_checkin := c.Params("id_checkin")

			checkin, err := checkinValidateService.ValidateCheckin(id_checkin)
			if err != nil {
				return handlers.HandleHTTPError(c, err)
			}

			return c.JSON(checkin)
		})

	app.Get("/checkin/history/:id_user", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_user := c.Params("id_user")

		count, err := checkinCountHistoryService.CountCheckinHistory(id_user)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(count)
	})

	app.Get("/checkin/history/:id_user", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_user := c.Params("id_user")
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid page",
				Code:    400,
			})
		}

		checkins, err := checkinListHistoryService.ListCheckinHistory(id_user, page)

		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(checkins)
	})
}
