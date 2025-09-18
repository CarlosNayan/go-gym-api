package services

import (
	"api-gym-on-go/src/config/errors"
	users_types "api-gym-on-go/src/modules/users/types"
)

type UsersMeService struct {
	UserRepository users_types.UserRepository
}

func NewUsersMeService(userRepo users_types.UserRepository) *UsersMeService {
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
		"id_user":   user.ID,
		"user_name": user.UserName,
		"email":     user.Email,
		"role":      string(user.Role),
	}

	return userData, nil
}
