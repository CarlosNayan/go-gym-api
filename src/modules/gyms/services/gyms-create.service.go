package services

import (
	"api-gym-on-go/src/models"
	gyms_schemas "api-gym-on-go/src/modules/gyms/schemas"
	gyms_types "api-gym-on-go/src/modules/gyms/types"
)

type GymsCreateService struct {
	GymRepository gyms_types.GymsRepository
}

func NewGymsCreateService(gymsRepo gyms_types.GymsRepository) *GymsCreateService {
	return &GymsCreateService{GymRepository: gymsRepo}
}

func (gcs *GymsCreateService) CreateGym(gym *gyms_schemas.GymCreateBody) (*models.Gym, error) {
	return gcs.GymRepository.CreateGym(gym)
}
