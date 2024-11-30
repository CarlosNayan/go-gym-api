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
			var params schemas.CheckinValidateParams

			if err := c.ParamsParser(&params); err != nil {
				fmt.Println(err)
				return handlers.HandleHTTPError(c, &errors.CustomError{
					Message: "Invalid request params",
					Code:    400,
				})
			}

			checkin, err := checkinValidateService.ValidateCheckin(params.IDCheckin)
			if err != nil {
				return handlers.HandleHTTPError(c, err)
			}

			return c.JSON(checkin)
		})

	app.Get("/checkin/history/count", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_user := c.Locals("sub").(string)

		count, err := checkinCountHistoryService.CountCheckinHistory(id_user)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(map[string]interface{}{"count": count})
	})

	app.Get("/checkin/history", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		id_user := c.Locals("sub").(string)
		var params schemas.CheckinValidateQuery

		if err := c.QueryParser(&params); err != nil {
			return handlers.HandleHTTPError(c, &errors.CustomError{
				Message: "Invalid page",
				Code:    400,
			})
		}

		page_num, err := strconv.Atoi(params.Page)

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
