package checkins

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/handlers"
	"api-gym-on-go/src/config/middleware"
	"api-gym-on-go/src/config/validate"
	"api-gym-on-go/src/modules/checkins/repository"
	checkin_schemas "api-gym-on-go/src/modules/checkins/schemas"
	"api-gym-on-go/src/modules/checkins/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, db *sql.DB) {
	app.Post("/checkin/create", middleware.ValidateJWT, func(ctx *fiber.Ctx) error {
		sub := ctx.Locals("sub").(string)

		body, err := validate.ParseBody[checkin_schemas.CheckinCreateBody](ctx)
		if err != nil {
			return handlers.HandleHTTPError(ctx, &errors.CustomError{Message: err.Error(), Code: 400})
		}

		repo := repository.NewCheckinRepository(db)
		service := services.NewCheckinCreateService(repo)

		err = service.Execute(sub, body)
		if err != nil {

			return handlers.HandleHTTPError(ctx, err)
		}

		return ctx.SendStatus(201)
	})

	app.Put("/checkin/validate/:id_checkin",
		middleware.ValidateJWT,
		middleware.VerifyUserRole("ADMIN"),
		func(ctx *fiber.Ctx) error {
			params, err := validate.ParseParams[checkin_schemas.CheckinValidateParams](ctx)
			if err != nil {
				return handlers.HandleHTTPError(ctx, &errors.CustomError{Message: err.Error(), Code: 400})
			}

			repo := repository.NewCheckinRepository(db)
			service := services.NewCheckinValidateService(repo)

			checkin, err := service.ValidateCheckin(params.IDCheckin)
			if err != nil {
				return handlers.HandleHTTPError(ctx, err)
			}

			return ctx.JSON(checkin)
		})

	app.Get("/checkin/history/count", middleware.ValidateJWT, func(c *fiber.Ctx) error {
		sub := c.Locals("sub").(string)

		repo := repository.NewCheckinRepository(db)
		service := services.NewCheckinCountHistory(repo)

		count, err := service.CountCheckinHistory(sub)
		if err != nil {
			return handlers.HandleHTTPError(c, err)
		}

		return c.JSON(map[string]interface{}{"count": count})
	})

	app.Get("/checkin/history", middleware.ValidateJWT, func(ctx *fiber.Ctx) error {
		sub := ctx.Locals("sub").(string)

		query, err := validate.ParseQueryParams[checkin_schemas.CheckinValidateQuery](ctx)
		if err != nil {
			return handlers.HandleHTTPError(ctx, &errors.CustomError{Message: err.Error(), Code: 400})
		}

		repo := repository.NewCheckinRepository(db)
		service := services.NewCheckinListHistory(repo)

		checkins, err := service.ListCheckinHistory(sub, query.Page)
		if err != nil {
			return handlers.HandleHTTPError(ctx, err)
		}

		return ctx.JSON(checkins)
	})

}
