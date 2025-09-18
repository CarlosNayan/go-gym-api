package users_types

import (
	"api-gym-on-go/src/models"
	users_schemas "api-gym-on-go/src/modules/users/schemas"
)

type UserRepository interface {
	GetProfileById(id string) (*models.User, error)
	UserEmailVerify(email string) (*string, error)
	CreateUser(data *users_schemas.UserCreateBody) (*models.User, error)
}
