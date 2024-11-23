package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/checkins/repository"
)

type CheckinListHistory struct {
	CheckinRepository *repository.CheckinRepository
}

func NewCheckinListHistory(checkinRepository *repository.CheckinRepository) *CheckinListHistory {
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
