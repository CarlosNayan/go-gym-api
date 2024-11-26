package gyms_service_test

import (
	"api-gym-on-go/src/modules/gyms/services"
	"api-gym-on-go/tests/services/gyms/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	gymsNearbyInMemoryRepository *repository.InMemoryGymsRepository
	gymsNearbyService            *services.GymsNearbyService
)

func setupNearbyService() {
	gymsNearbyInMemoryRepository = repository.NewInMemoryGymRepository()
	gymsNearbyService = services.NewGymsNearbyService(gymsNearbyInMemoryRepository)
}

func TestGymsNearby(t *testing.T) {
	t.Run("should be able to register", func(t *testing.T) {
		setupNearbyService()

		gyms, err := gymsNearbyService.GetGymsNearby(1.244567, 1.244567)

		require.Len(t, gyms, 1)
		require.NoError(t, err)
	})

	t.Run("should not be able to register with invalid coordinates", func(t *testing.T) {
		setupNearbyService()

		_, err := gymsNearbyService.GetGymsNearby(0, 0)

		require.Error(t, err)
	})
}
