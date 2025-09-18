package services

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/models"
	gyms_types "api-gym-on-go/src/modules/gyms/types"
)

type GymsNearbyService struct {
	GymRepository gyms_types.GymsRepository
}

func NewGymsNearbyService(gymsRepo gyms_types.GymsRepository) *GymsNearbyService {
	return &GymsNearbyService{GymRepository: gymsRepo}
}

func (gns *GymsNearbyService) GetGymsNearby(latitude, longitude float64) ([]models.Gym, error) {
	if latitude == 0 || longitude == 0 {
		return nil, &errors.InvalidCoordinatesError{}
	}

	return gns.GymRepository.GymsNearby(latitude, longitude)
}
