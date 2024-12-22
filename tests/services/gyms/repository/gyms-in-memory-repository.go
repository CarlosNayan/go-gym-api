package repository

import (
	"api-gym-on-go/src/config/utils"
	"api-gym-on-go/src/models"
	"api-gym-on-go/src/modules/gyms/interfaces"
	"strings"

	"github.com/google/uuid"
)

type InMemoryGymsRepository struct {
	items []models.Gym
}

func NewInMemoryGymRepository() *InMemoryGymsRepository {
	return &InMemoryGymsRepository{
		items: []models.Gym{
			{
				ID:          "0ebd4f88-d712-4b0f-9278-41d595c690ad",
				GymName:     "Default Gym",
				Description: nil,
				Phone:       nil,
				Latitude:    1.234567,
				Longitude:   1.234567,
			},
		},
	}
}

var _ interfaces.GymsRepository = (*InMemoryGymsRepository)(nil)

func (i *InMemoryGymsRepository) CreateGym(gym *models.Gym) error {
	newGym := models.Gym{
		ID:          uuid.New().String(),
		GymName:     gym.GymName,
		Description: gym.Description,
		Phone:       gym.Phone,
		Latitude:    gym.Latitude,
		Longitude:   gym.Longitude,
	}

	i.items = append(i.items, newGym)
	return nil
}

func (i *InMemoryGymsRepository) SearchGyms(query string) ([]models.Gym, error) {
	var result []models.Gym

	for _, gym := range i.items {
		if strings.Contains(gym.GymName, query) {
			result = append(result, gym)
		}
	}

	return result, nil
}

func (i *InMemoryGymsRepository) GymsNearby(latitude float64, longitude float64) ([]models.Gym, error) {
	var result []models.Gym

	for _, gym := range i.items {
		from := utils.Coordinate{Latitude: latitude, Longitude: longitude}
		to := utils.Coordinate{Latitude: gym.Latitude, Longitude: gym.Longitude}

		distance := utils.GetDistanceBetweenCoordinates(from, to)

		if distance <= 10 {
			result = append(result, gym)
		}
	}

	return result, nil
}
