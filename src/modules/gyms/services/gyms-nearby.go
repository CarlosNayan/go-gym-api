package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/modules/gyms/interfaces"
)

type GymsNearbyService struct {
	GymRepository interfaces.GymsRepository
}

func NewGymsNearbyService(gymsRepo interfaces.GymsRepository) *GymsNearbyService {
	return &GymsNearbyService{GymRepository: gymsRepo}
}

func (gns *GymsNearbyService) GetGymsNearby(latitude, longitude float64) ([]models.Gym, error) {
	if latitude == 0 || longitude == 0 {
		return nil, &errors.InvalidCoordinatesError{}
	}

	return gns.GymRepository.GymsNearby(latitude, longitude)
}
