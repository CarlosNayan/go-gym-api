package services

import (
	"api-gym-on-go/src/modules/checkins/repository"
)

type CheckinCountHistory struct {
	CheckinRepository *repository.CheckinRepository
}

func NewCheckinCountHistory(checkinRepository *repository.CheckinRepository) *CheckinCountHistory {
	return &CheckinCountHistory{CheckinRepository: checkinRepository}
}

func (cch *CheckinCountHistory) CountCheckinHistory(id_user string) (int64, error) {
	return cch.CheckinRepository.CountByUserId(id_user)
}
