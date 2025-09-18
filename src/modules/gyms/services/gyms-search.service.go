package services

import (
	"api-gym-on-go/src/models"
	gyms_types "api-gym-on-go/src/modules/gyms/types"
)

type GymsSearchService struct {
	GymRepository gyms_types.GymsRepository
}

func NewGymsSearchService(gymsRepo gyms_types.GymsRepository) *GymsSearchService {
	return &GymsSearchService{GymRepository: gymsRepo}
}

func (gss *GymsSearchService) SearchGyms(query string) ([]models.Gym, error) {
	return gss.GymRepository.SearchGyms(query)
}
