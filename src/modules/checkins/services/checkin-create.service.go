package services

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/config/utils"
	"api-gym-on-go/src/models"
	checkin_schemas "api-gym-on-go/src/modules/checkins/schemas"
	checkin_types "api-gym-on-go/src/modules/checkins/types"
)

type CheckinCreate struct {
	checkinsRepository checkin_types.CheckinsRepository
}

func NewCheckinCreateService(checkinsRepository checkin_types.CheckinsRepository) *CheckinCreate {
	return &CheckinCreate{checkinsRepository: checkinsRepository}
}

func (cc *CheckinCreate) Execute(IDUser string, body *checkin_schemas.CheckinCreateBody) error {
	checkinAlrightExistsToday, err := cc.checkinsRepository.FindCheckinByIdOnDate(IDUser)
	if err != nil {
		return err
	} else if checkinAlrightExistsToday != nil && checkinAlrightExistsToday.IDUser == IDUser {
		return &errors.MaxNumberOfCheckinsError{}
	}

	gym, err := cc.checkinsRepository.FindGymByID(body.IDGym)
	if err != nil {
		return err
	} else if gym == nil {
		return &errors.ResourceNotFoundError{}
	}

	from := utils.Coordinate{Latitude: body.UserLatitude, Longitude: body.UserLongitude}
	to := utils.Coordinate{Latitude: gym.Latitude, Longitude: gym.Longitude}

	distance := utils.GetDistanceBetweenCoordinates(from, to)

	if distance > 1 {
		return &errors.InvalidCoordinatesError{}
	}

	checkin := models.Checkin{
		IDUser: IDUser,
		IDGym:  body.IDGym,
	}

	return cc.checkinsRepository.CreateCheckin(&checkin)
}
