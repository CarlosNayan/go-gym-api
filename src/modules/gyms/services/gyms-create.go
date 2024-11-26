package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/gyms/interfaces"
)

type GymsCreateService struct {
	GymRepository interfaces.GymsRepository
}

func NewGymsCreateService(gymsRepo interfaces.GymsRepository) *GymsCreateService {
	return &GymsCreateService{GymRepository: gymsRepo}
}

func (gcs *GymsCreateService) CreateGym(gym *models.Gym) error {
	return gcs.GymRepository.CreateGym(gym)
}
