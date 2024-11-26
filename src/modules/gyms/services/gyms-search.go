package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/gyms/interfaces"
)

type GymsSearchService struct {
	GymRepository interfaces.GymsRepository
}

func NewGymsSearchService(gymsRepo interfaces.GymsRepository) *GymsSearchService {
	return &GymsSearchService{GymRepository: gymsRepo}
}

func (gss *GymsSearchService) SearchGyms(query string) ([]models.Gym, error) {
	return gss.GymRepository.SearchGyms(query)
}
