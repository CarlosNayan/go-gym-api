package checkins

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/handlers"
	"api-gym-on-go/src/config/middleware"
	"api-gym-on-go/src/modules/checkins/repository"
	"api-gym-on-go/src/modules/checkins/schemas"
	"api-gym-on-go/src/modules/checkins/services"
	"fmt"
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
		var body schemas.CheckinCreateBody
		IDUser := c.Locals("sub").(string)

		if err := c.BodyParser(&body); err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid request body",
				Code:    400,
			})
		}

		body.IDUser = IDUser

		if err := body.Validate(); err != nil {

			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: fmt.Sprintf("Validation failed: %v", err),
				Code:    400,
			})
		}

		err := checkinCreateService.CreateCheckin(&body)
		if err != nil {

			return handlers.HandleHTTPError(c, err)
		}

		return c.SendStatus(201)
	})

	app.Put("/checkin/validate/:id_checkin",
		middleware.ValidateJWT,
		middleware.VerifyUserRole("ADMIN"),
		func(c *fiber.Ctx) error {
			var IdCheckin string

			if err := c.ParamsParser(&IdCheckin); err != nil {
				return handlers.HandleHTTPError(c, &errors.CustomError{
					Message: "Invalid request params",
					Code:    400,
				})
			}

			checkin, err := checkinValidateService.ValidateCheckin(IdCheckin)
			if err != nil {
				return handlers.HandleHTTPError(c, err)
			}

			return c.JSON(checkin)
		})

	app.Get("/checkin/history/:id_user", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		var IDUser string
		if err := c.ParamsParser(&IDUser); err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid request params",
				Code:    400,
			})
		}

		count, err := checkinCountHistoryService.CountCheckinHistory(IDUser)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(count)
	})

	app.Get("/checkin/history/:id_user", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		var id_user string
		var page string

		if err := c.ParamsParser(&id_user); err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid request params",
				Code:    400,
			})
		}

		if err := c.QueryParser(&page); err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid page",
				Code:    400,
			})
		}

		page_num, err := strconv.Atoi(page)
		if err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid page number",
				Code:    400,
			})
		}

		checkins, err := checkinListHistoryService.ListCheckinHistory(id_user, page_num)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(checkins)
	})

}
