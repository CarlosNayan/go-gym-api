package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/gyms/repository"
)

type GymsCreateService struct {
	GymRepository *repository.GymsRepository
}

func NewGymsCreateService(gymsRepo *repository.GymsRepository) *GymsCreateService {
	return &GymsCreateService{GymRepository: gymsRepo}
}

func (gcs *GymsCreateService) CreateGym(gym *models.Gym) error {
	return gcs.GymRepository.CreateGym(gym)
}
