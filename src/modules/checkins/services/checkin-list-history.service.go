package services

import (
	"api-gym-on-go/src/models"
	checkin_types "api-gym-on-go/src/modules/checkins/types"
)

type CheckinListHistory struct {
	CheckinRepository checkin_types.CheckinsRepository
}

func NewCheckinListHistory(checkinRepository checkin_types.CheckinsRepository) *CheckinListHistory {
	return &CheckinListHistory{CheckinRepository: checkinRepository}
}

func (clh *CheckinListHistory) ListCheckinHistory(id_user string, page int) ([]models.Checkin, error) {
	var checkins []models.Checkin
	checkins, err := clh.CheckinRepository.ListAllCheckinsHistoryOfUser(id_user, page)
	if err != nil {
		return nil, err
	}
	return checkins, nil
}
