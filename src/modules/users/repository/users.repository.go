package repository

import (
	"gorm.io/gorm"

	"api-gym-on-go/schema"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByID(id string) (*schema.User, error) {
	var user schema.User

	result := r.DB.Where("id_user = ?", id).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) Create(user *schema.User) error {
	return r.DB.Create(user).Error
}
