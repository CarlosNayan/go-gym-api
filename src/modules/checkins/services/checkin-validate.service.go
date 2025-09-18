package services

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/utils"
	"api-gym-on-go/src/models"
	checkin_types "api-gym-on-go/src/modules/checkins/types"
)

type CheckinValidate struct {
	CheckinRepository checkin_types.CheckinsRepository
}

func NewCheckinValidateService(checkinRepository checkin_types.CheckinsRepository) *CheckinValidate {
	return &CheckinValidate{CheckinRepository: checkinRepository}
}

func (cv *CheckinValidate) ValidateCheckin(id_checkin string) (nill *models.Checkin, err error) {
	checkin, err := cv.CheckinRepository.FindCheckinById(id_checkin)
	if err != nil {
		return nil, err
	}
	if checkin == nil {
		return nil, &errors.ResourceNotFoundError{}
	}
	if checkin.ValidatedAt != nil {
		return nil, &errors.CustomError{Message: "check-in already validated", Code: 400}
	}

	timeNow, _ := utils.NewMoment()
	checkinCreatedAtMoment, _ := utils.NewMoment(checkin.CreatedAt)

	difference := timeNow.Diff(checkinCreatedAtMoment, "minutes")

	if difference > 20 {
		return nil, &errors.CustomError{Message: "check-in expired", Code: 400}
	}

	validatedCheckin, err := cv.CheckinRepository.UpdateCheckin(id_checkin)
	if err != nil {
		return nil, err
	}

	return validatedCheckin, nil
}
