package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/modules/gyms/repository"
)

type GymsNearbyService struct {
	GymRepository *repository.GymsRepository
}

func NewGymsNearbyService(gymsRepo *repository.GymsRepository) *GymsNearbyService {
	return &GymsNearbyService{GymRepository: gymsRepo}
}

func (gns *GymsNearbyService) GetGymsNearby(latitude, longitude float64) ([]models.Gym, error) {
	if latitude == 0 || longitude == 0 {
		return nil, &errors.InvalidCoordinatesError{}
	}

	return gns.GymRepository.SearchGymsNearby(latitude, longitude)
}
