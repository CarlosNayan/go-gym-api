package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/utils"
	"api-gym-on-go/src/modules/checkins/repository"
	"errors"
	"time"
)

type CheckinValidate struct {
	CheckinRepository *repository.CheckinRepository
}

func NewCheckinValidateService(checkinRepository *repository.CheckinRepository) *CheckinValidate {
	return &CheckinValidate{CheckinRepository: checkinRepository}
}

func (cv *CheckinValidate) ValidateCheckin(id_checkin string) (nill *models.Checkin, err error) {
	checkin, err := cv.CheckinRepository.FindCheckinById(id_checkin)
	if err != nil {
		return nil, err
	}

	moment, err := utils.NewMoment(checkin.CreatedAt)
	if err != nil {
		return nil, err
	}

	if moment.Diff(utils.Time(time.Now()), "minutes") > 20 {
		return nil, errors.New("the check-in can only be validated until 20 minutes of its creation")
	}

	validatedCheckin, err := cv.CheckinRepository.UpdateCheckin(id_checkin)
	if err != nil {
		return nil, err
	}

	return validatedCheckin, nil
}
