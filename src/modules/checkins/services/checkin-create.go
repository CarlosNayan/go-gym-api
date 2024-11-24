package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/checkins/repository"
)

type CheckinCreate struct {
	checkinsRepository *repository.CheckinRepository
}

func NewCheckinCreateService(checkinsRepository *repository.CheckinRepository) *CheckinCreate {
	return &CheckinCreate{checkinsRepository: checkinsRepository}
}

func (cc *CheckinCreate) CreateCheckin(IDUser string, checkin *models.Checkin) error {
	checkin.IDUser = IDUser
	return cc.checkinsRepository.CreateCheckin(checkin)
}
