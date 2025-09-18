package repository

import (
	"api-gym-on-go/src/config/utils"
	"api-gym-on-go/src/models"
	gyms_schemas "api-gym-on-go/src/modules/gyms/schemas"
	gyms_types "api-gym-on-go/src/modules/gyms/types"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type GymsRepository struct {
	DB *sql.DB
}

func NewGymsRepository(db *sql.DB) *GymsRepository {
	return &GymsRepository{db}
}

var _ gyms_types.GymsRepository = &GymsRepository{}

func (gr *GymsRepository) CreateGym(gym *gyms_schemas.GymCreateBody) (*models.Gym, error) {
	id := uuid.New()

	var result models.Gym

	query := `
		INSERT INTO gyms
		(id_gym, gym_name, description, latitude, longitude, phone)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING id_gym, gym_name, description, latitude, longitude, phone
	`

	err := gr.DB.
		QueryRow(query, id, gym.GymName, gym.Description, gym.Latitude, gym.Longitude, gym.Phone).
		Scan(&result.ID, &result.GymName, &result.Description, &result.Latitude, &result.Longitude, &result.Phone)
	if err != nil {
		return nil, fmt.Errorf("error inserting checkin: %w", err)
	}

	return &result, nil
}

func (gr *GymsRepository) GymsNearby(latitude, longitude float64) ([]models.Gym, error) {
	var gyms []models.Gym

	query := `
		SELECT *
		FROM gyms
		WHERE (6371 * acos(cos(radians($1)) * cos(radians(latitude)) * cos(radians(longitude) - radians($2)) + sin(radians($1)) * sin(radians(latitude)))) <= 10
	`

	rows, err := gr.DB.Query(query, latitude, longitude)
	if err != nil {
		return nil, utils.WrapError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var gym models.Gym
		err = rows.Scan(&gym.ID, &gym.GymName, &gym.Description, &gym.Phone, &gym.Latitude, &gym.Longitude)
		if err != nil {
			return nil, utils.WrapError(err)
		}
		gyms = append(gyms, gym)
	}

	if err = rows.Err(); err != nil {
		return nil, utils.WrapError(err)
	}

	if len(gyms) == 0 {
		return []models.Gym{}, nil
	}

	return gyms, nil
}

func (gr *GymsRepository) SearchGyms(searchQuery string) ([]models.Gym, error) {
	var gyms []models.Gym

	query := `
		SELECT * FROM gyms
		WHERE gym_name LIKE $1
	`

	rows, err := gr.DB.Query(query, "%"+searchQuery+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var gym models.Gym
		err = rows.Scan(&gym.ID, &gym.GymName, &gym.Description, &gym.Phone, &gym.Latitude, &gym.Longitude)
		if err != nil {
			return nil, utils.WrapError(err)
		}
		gyms = append(gyms, gym)
	}

	if err = rows.Err(); err != nil {
		return nil, utils.WrapError(err)
	}

	if len(gyms) == 0 {
		return []models.Gym{}, nil
	}

	return gyms, nil
}
