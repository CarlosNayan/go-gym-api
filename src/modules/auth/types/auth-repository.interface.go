package auth_types

import (
	"api-gym-on-go/src/models"
)

type IAuthRepository interface {
	FindByEmail(email string) (*models.User, error)
}
