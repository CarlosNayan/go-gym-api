package services

import (
	"api-gym-on-go/schema"
	"api-gym-on-go/src/modules/users/repository"
)

type UsersCreateService struct {
	UserRepository *repository.UserRepository
}

func NewUsersCreateService(userRepo *repository.UserRepository) *UsersCreateService {
	return &UsersCreateService{UserRepository: userRepo}
}

func (ucs *UsersCreateService) CreateUser(user *schema.User) error {
	return ucs.UserRepository.Create(user)
}
