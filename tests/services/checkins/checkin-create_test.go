package checkins_test

import (
	"api-gym-on-go/src/modules/checkins/schemas"
	"api-gym-on-go/src/modules/checkins/services"
	"api-gym-on-go/tests/services/checkins/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	checkinsCreateInMemoryRepository *repository.InMemoryCheckinsRepository
	checkinsCreateService            *services.CheckinCreate
)

func setupCreateService() {
	checkinsCreateInMemoryRepository = &repository.InMemoryCheckinsRepository{}
	checkinsCreateService = services.NewCheckinCreateService(checkinsCreateInMemoryRepository)
}

func TestCheckinCreate(t *testing.T) {

	var checkin schemas.CheckinCreateBody

	checkin.IDUser = "1e2d4f88-d712-4b0f-9278-41d595c690ad"
	checkin.IDGym = "2e2d4f88-d712-4b0f-9278-41d595c690ad"

	t.Run("should be able to create a checkin", func(t *testing.T) {
		setupCreateService()

		err := checkinsCreateService.CreateCheckin(&checkin)

		require.NoError(t, err)
	})
}
