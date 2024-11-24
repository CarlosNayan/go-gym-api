package interfaces

import "api-gym-on-go/models"

type UserRepository interface {
	GetProfileById(id string) (*models.User, error)
	UserEmailVerify(email string) (string, error)
	CreateUser(data *models.User) (*models.User, error)
}
