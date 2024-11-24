package repository

import (
	"gorm.io/gorm"

	"api-gym-on-go/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetProfileById(id string) (*models.User, error) {
	var user models.User

	result := r.DB.Where("id_user = ?", id).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) UserEmailVerify(email string) (string, error) {
	var user models.User

	result := r.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", result.Error
	}

	if user.Email == email {
		return user.Email, nil
	} else {
		return "", nil
	}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	var createdUser models.User
	result := r.DB.Create(user).First(&createdUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &createdUser, nil
}
