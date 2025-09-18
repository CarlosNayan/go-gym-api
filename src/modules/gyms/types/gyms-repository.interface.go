package gyms_types

import (
	"api-gym-on-go/src/models"
	gyms_schemas "api-gym-on-go/src/modules/gyms/schemas"
)

type GymsRepository interface {
	CreateGym(gym *gyms_schemas.GymCreateBody) (*models.Gym, error)
	GymsNearby(latitude, longitude float64) ([]models.Gym, error)
	SearchGyms(query string) ([]models.Gym, error)
}
