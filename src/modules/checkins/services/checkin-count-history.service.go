package services

import checkin_types "api-gym-on-go/src/modules/checkins/types"

type CheckinCountHistory struct {
	CheckinRepository checkin_types.CheckinsRepository
}

func NewCheckinCountHistory(checkinRepository checkin_types.CheckinsRepository) *CheckinCountHistory {
	return &CheckinCountHistory{CheckinRepository: checkinRepository}
}

func (cch *CheckinCountHistory) CountCheckinHistory(id_user string) (int64, error) {
	return cch.CheckinRepository.CountByUserId(id_user)
}
