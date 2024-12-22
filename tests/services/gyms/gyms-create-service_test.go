package gyms_service_test

import (
	"api-gym-on-go/src/models"
	"api-gym-on-go/src/modules/gyms/services"
	"api-gym-on-go/tests/services/gyms/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	gymsCreateInMemoryRepository *repository.InMemoryGymsRepository
	gymsCreateService            *services.GymsCreateService
)

func setupCreateService() {
	gymsCreateInMemoryRepository = &repository.InMemoryGymsRepository{}
	gymsCreateService = services.NewGymsCreateService(gymsCreateInMemoryRepository)
}

func TestGymCreate(t *testing.T) {
	t.Run("should be able to register a gym", func(t *testing.T) {
		setupCreateService()

		description := "Test description"
		phone := "1234567890"

		err := gymsCreateService.CreateGym(&models.Gym{
			GymName:     "Test gym",
			Description: &description,
			Phone:       &phone,
			Latitude:    0,
			Longitude:   0,
		})

		require.NoError(t, err)
	})
}
