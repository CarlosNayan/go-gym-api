package gyms_service_test

import (
	"api-gym-on-go/src/modules/gyms/services"
	"api-gym-on-go/tests/services/gyms/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	gymsSearchInMemoryRepository *repository.InMemoryGymsRepository
	gymsSearchService            *services.GymsSearchService
)

func setupSearchService() {
	gymsSearchInMemoryRepository = repository.NewInMemoryGymRepository()
	gymsSearchService = services.NewGymsSearchService(gymsSearchInMemoryRepository)
}

func TestGymsSearch(t *testing.T) {
	t.Run("should be able to search", func(t *testing.T) {
		setupSearchService()

		gyms, err := gymsSearchService.SearchGyms("Default")

		require.Len(t, gyms, 1)
		require.NoError(t, err)
	})

	t.Run("should not be able to search with invalid name", func(t *testing.T) {
		setupSearchService()

		gyms, err := gymsSearchService.SearchGyms("Invalid")

		require.Len(t, gyms, 0)
		require.NoError(t, err)
	})
}
