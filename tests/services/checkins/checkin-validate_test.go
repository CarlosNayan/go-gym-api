package checkins_test

import (
	"api-gym-on-go/src/modules/checkins/services"
	"api-gym-on-go/tests/services/checkins/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	checkinsValidateInMemoryRepository *repository.InMemoryCheckinsRepository
	checkinsValidateService            *services.CheckinValidate
)

func setupValidateService() {
	checkinsValidateInMemoryRepository = repository.NewInMemoryCheckinsRepository()
	checkinsValidateService = services.NewCheckinValidateService(checkinsValidateInMemoryRepository)
}

func TestCheckinValidate(t *testing.T) {
	t.Run("should be able to validate a checkin", func(t *testing.T) {
		setupValidateService()

		checkins, err := checkinsValidateService.ValidateCheckin("0ebd4f88-d712-4b0f-9278-41d595c690ad")

		require.NotNil(t, checkins.ValidatedAt, "ValidatedAt should not be nil")
		require.NoError(t, err)
	})
}
