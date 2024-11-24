package services

import (
	"api-gym-on-go/src/config/errors"
	"api-gym-on-go/src/modules/users/interfaces"
)

type UsersMeService struct {
	UserRepository interfaces.UserRepository
}

func NewUsersMeService(userRepo interfaces.UserRepository) *UsersMeService {
	return &UsersMeService{UserRepository: userRepo}
}

func (ums *UsersMeService) GetUserByID(id string) (map[string]string, error) {
	user, err := ums.UserRepository.GetProfileById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, &errors.ResourceNotFoundError{}
	}

	userData := map[string]string{
		"id":       user.ID,
		"userName": user.UserName,
		"email":    user.Email,
		"role":     string(user.Role),
	}

	return userData, nil
}
