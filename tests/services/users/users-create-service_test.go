package users_service_test

import (
	"testing"

	"api-gym-on-go/models"
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/modules/users/services"
	"api-gym-on-go/tests/services/users/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

var (
	usersInMemoryRepository *repository.InMemoryUserRepository
	userService             *services.UsersCreateService
)

func setupCreateService() {
	usersInMemoryRepository = &repository.InMemoryUserRepository{}
	userService = services.NewUsersCreateService(usersInMemoryRepository)
}

func TestUserCreate(t *testing.T) {
	t.Run("should be able to register", func(t *testing.T) {
		setupCreateService()

		user, err := userService.CreateUser(&models.User{
			UserName: "Jhon Doe",
			Email:    "email@email.com",
			Password: "123456",
		})

		require.NoError(t, err)
		assert.NotNil(t, user.ID)
	})

	t.Run("should hash user password upon registration", func(t *testing.T) {
		setupCreateService()

		user, err := userService.CreateUser(&models.User{
			UserName: "Jhon Doe",
			Email:    "email@email.com",
			Password: "123456",
		})

		require.NoError(t, err)

		// Verifica se a senha foi criptografada corretamente
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("123456"))
		assert.NoError(t, err)
	})

	t.Run("should not be able to register with same email twice", func(t *testing.T) {
		setupCreateService()

		email := "email@email.com"

		_, err := userService.CreateUser(&models.User{
			UserName: "Jhon Doe",
			Email:    email,
			Password: "123456",
		})
		require.NoError(t, err)

		// Tenta registrar com o mesmo email
		_, err = userService.CreateUser(&models.User{
			UserName: "Jhon Doe",
			Email:    email,
			Password: "123456",
		})

		// Espera o erro de usuário já existente
		var userAlreadyExistsErr *errors.UserAlreadyExistsError
		assert.ErrorAs(t, err, &userAlreadyExistsErr)
	})
}
