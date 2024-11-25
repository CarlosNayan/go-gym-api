package schemas

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CheckinCreateBody struct {
	IDUser        string  `json:"id_user" validate:"required"`
	IDGym         string  `json:"id_gym" validate:"required"`
	UserLatitude  float64 `json:"user_latitude" validate:"required,latitude"`
	UserLongitude float64 `json:"user_longitude" validate:"required,longitude"`
}

func latitude(fl validator.FieldLevel) bool {
	lat := fl.Field().Float()
	return lat >= -90 && lat <= 90
}

func longitude(fl validator.FieldLevel) bool {
	long := fl.Field().Float()
	return long >= -180 && long <= 180
}

func (c *CheckinCreateBody) Validate() map[string]string {
	validate := validator.New()

	validate.RegisterValidation("latitude", latitude)
	validate.RegisterValidation("longitude", longitude)

	err := validate.Struct(c)
	if err != nil {

		errors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {

			switch err.Tag() {
			case "required":
				errors[err.Field()] = fmt.Sprintf("%s é obrigatório", err.Field())
			case "latitude":
				errors[err.Field()] = "Latitude deve estar entre -90 e 90"
			case "longitude":
				errors[err.Field()] = "Longitude deve estar entre -180 e 180"
			default:
				errors[err.Field()] = fmt.Sprintf("%s é inválido", err.Field())
			}
		}

		return errors
	}

	return nil
}
