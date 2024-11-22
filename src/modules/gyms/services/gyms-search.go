package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/gyms/repository"
)

type GymsSearchService struct {
	GymRepository *repository.GymsRepository
}

func NewGymsSearchService(gymsRepo *repository.GymsRepository) *GymsSearchService {
	return &GymsSearchService{GymRepository: gymsRepo}
}

func (gss *GymsSearchService) SearchGyms(query string) ([]models.Gym, error) {
	return gss.GymRepository.SearchGyms(query)
}
