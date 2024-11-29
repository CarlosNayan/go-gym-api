package checkins_test

import (
	"api-gym-on-go/src/modules/checkins/services"
	"api-gym-on-go/tests/services/checkins/repository"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	checkinsListHistoryInMemoryRepository *repository.InMemoryCheckinsRepository
	checkinsListHistoryService            *services.CheckinListHistory
)

func setupListHistoryService() {
	checkinsListHistoryInMemoryRepository = repository.NewInMemoryCheckinsRepository()
	checkinsListHistoryService = services.NewCheckinListHistory(checkinsListHistoryInMemoryRepository)
}

func TestCheckinListHistory(t *testing.T) {
	t.Run("should be able to list history", func(t *testing.T) {
		setupListHistoryService()

		checkins, err := checkinsListHistoryService.ListCheckinHistory("1e2d4f88-d712-4b0f-9278-41d595c690ad", 1)

		fmt.Println(checkins)

		require.Equal(t, len(checkins), 1)
		require.NoError(t, err)
	})
}
