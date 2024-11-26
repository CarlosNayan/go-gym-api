package users_service_test

import (
	"api-gym-on-go/src/modules/users/services"
	"api-gym-on-go/tests/services/users/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	usersMeInMemoryRepository *repository.InMemoryUserRepository
	userMeService             *services.UsersMeService
)

func setupUserMe() {
	usersMeInMemoryRepository = repository.NewInMemoryUserRepository()
	userMeService = services.NewUsersMeService(usersMeInMemoryRepository)
}

func TestUserMe(t *testing.T) {
	t.Run("should be able to get profile", func(t *testing.T) {
		setupUserMe()

		user, err := userMeService.GetUserByID("0ebd4f88-d712-4b0f-9278-41d595c690ad")

		require.NoError(t, err)
		assert.NotNil(t, user["id_user"])
	})
}
