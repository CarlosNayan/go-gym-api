package services

import (
	"api-gym-on-go/src/modules/checkins/interfaces"
)

type CheckinCountHistory struct {
	CheckinRepository interfaces.CheckinsRepository
}

func NewCheckinCountHistory(checkinRepository interfaces.CheckinsRepository) *CheckinCountHistory {
	return &CheckinCountHistory{CheckinRepository: checkinRepository}
}

func (cch *CheckinCountHistory) CountCheckinHistory(id_user string) (int64, error) {
	return cch.CheckinRepository.CountByUserId(id_user)
}
