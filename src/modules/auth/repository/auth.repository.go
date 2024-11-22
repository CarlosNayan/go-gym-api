package repository

import (
	"api-gym-on-go/schema"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(email string) (*schema.User, error) {
	var user schema.User

	result := r.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}
