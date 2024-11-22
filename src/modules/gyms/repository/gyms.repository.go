package repository

import (
	"api-gym-on-go/models"

	"gorm.io/gorm"
)

type GymsRepository struct {
	DB *gorm.DB
}

func NewGymsRepository(db *gorm.DB) *GymsRepository {
	return &GymsRepository{DB: db}
}

func (gr *GymsRepository) CreateGym(gym *models.Gym) error {
	return gr.DB.Create(gym).Error
}

func (gr *GymsRepository) SearchGymsNearby(latitude, longitude float64) ([]models.Gym, error) {
	var gyms []models.Gym

	query := `
		SELECT * FROM gyms
		WHERE (6371 * acos(cos(radians(?)) * cos(radians(latitude)) * cos(radians(longitude) - radians(?)) + sin(radians(?)) * sin(radians(latitude)))) <= 10
	`

	err := gr.DB.Raw(query, latitude, longitude, latitude).Scan(&gyms).Error
	if err != nil {
		return nil, err
	}

	return gyms, nil
}

func (gr *GymsRepository) SearchGyms(query string) ([]models.Gym, error) {
	var gyms []models.Gym

	err := gr.DB.Where("gym_name ILIKE ?", "%"+query+"%").Find(&gyms).Error
	if err != nil {
		return nil, err
	}

	return gyms, nil
}
