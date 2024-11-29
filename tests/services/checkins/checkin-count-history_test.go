package checkins_test

import (
	"api-gym-on-go/src/modules/checkins/services"
	"api-gym-on-go/tests/services/checkins/repository"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	checkinsCountHistoryInMemoryRepository *repository.InMemoryCheckinsRepository
	checkinsCountHistoryService            *services.CheckinCountHistory
)

func setupCountHistoryService() {
	checkinsCountHistoryInMemoryRepository = repository.NewInMemoryCheckinsRepository()
	checkinsCountHistoryService = services.NewCheckinCountHistory(checkinsCountHistoryInMemoryRepository)
}

func TestCheckinCountHistory(t *testing.T) {
	t.Run("should be able to count history", func(t *testing.T) {
		setupCountHistoryService()

		checkins, err := checkinsCountHistoryService.CountCheckinHistory("1e2d4f88-d712-4b0f-9278-41d595c690ad")

		fmt.Println(checkins)

		require.EqualValues(t, checkins, 1)
		require.NoError(t, err)
	})
}
