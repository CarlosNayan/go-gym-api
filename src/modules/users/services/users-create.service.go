package services

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/users/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UsersCreateService struct {
	UserRepository *repository.UserRepository
}

func NewUsersCreateService(userRepo *repository.UserRepository) *UsersCreateService {
	return &UsersCreateService{UserRepository: userRepo}
}

func (ucs *UsersCreateService) CreateUser(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("invalid password")
	}

	user.PasswordHash = string(hash)

	return ucs.UserRepository.Create(user)
}
