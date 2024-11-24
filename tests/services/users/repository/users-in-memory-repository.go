package repository

import (
	"errors"
	"time"

	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/users/interfaces"

	"github.com/google/uuid"
)

type InMemoryUserRepository struct {
	items []models.User
}

// Novo repositório em memória
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}

var _ interfaces.UserRepository = (*InMemoryUserRepository)(nil)

// Verifica se um usuário existe pelo ID
func (repo *InMemoryUserRepository) GetProfileById(id string) (*models.User, error) {
	for _, user := range repo.items {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (repo *InMemoryUserRepository) UserEmailVerify(email string) (string, error) {
	for _, user := range repo.items {
		if user.Email == email {
			return user.Email, nil
		}
	}
	return "", nil
}

// Cria um novo usuário
func (repo *InMemoryUserRepository) CreateUser(data *models.User) (*models.User, error) {
	user := models.User{
		ID:           uuid.New().String(),
		UserName:     data.UserName,
		Email:        data.Email,
		PasswordHash: data.PasswordHash,
		CreatedAt:    time.Now(),
	}

	repo.items = append(repo.items, user)
	return &user, nil
}
