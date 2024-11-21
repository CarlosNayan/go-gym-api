package services

import (
	"api-gym-on-go/src/modules/users/repository"
	"errors"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

func (s *UserService) GetUserByID(id string) (map[string]string, error) {
	user, err := s.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	userData := map[string]string{
		"id":       user.ID,
		"userName": user.UserName,
		"email":    user.Email,
		"role":     string(user.Role),
	}

	return userData, nil
}
