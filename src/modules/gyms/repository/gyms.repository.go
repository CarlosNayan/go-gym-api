package repository

import (
	"api-gym-on-go/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type GymsRepository struct {
	DB *sql.DB
}

func NewGymsRepository(db *sql.DB) *GymsRepository {
	return &GymsRepository{DB: db}
}

func (gr *GymsRepository) CreateGym(gym *models.Gym) error {
	id := uuid.New()

	query := `
		INSERT INTO gyms
		(id_gym, gym_name, description, latitude, longitude)
		VALUES
		($1, $2, $3, $4, $5)
	`

	_, err := gr.DB.Exec(query, id, gym.GymName, gym.Description, gym.Latitude, gym.Longitude)
	if err != nil {
		return fmt.Errorf("error inserting checkin: %w", err)
	}

	return nil
}

func (gr *GymsRepository) SearchGymsNearby(latitude, longitude float64) ([]models.Gym, error) {
	var gyms []models.Gym

	query := `
		SELECT id, gym_name, description, latitude, longitude FROM gyms
		WHERE (6371 * acos(cos(radians($1)) * cos(radians(latitude)) * cos(radians(longitude) - radians($2)) + sin(radians($1)) * sin(radians(latitude)))) <= 10
	`

	rows, err := gr.DB.Query(query, latitude, longitude, latitude)
	if err != nil {
		return nil, fmt.Errorf("error executing query to find nearby gyms: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var gym models.Gym
		err = rows.Scan(&gym.ID, &gym.GymName, &gym.Description, &gym.Latitude, &gym.Longitude)
		if err != nil {
			return nil, fmt.Errorf("error scanning gym row: %w", err)
		}
		gyms = append(gyms, gym)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
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
		err = rows.Scan(&gym.ID, &gym.GymName, &gym.Description, &gym.Latitude, &gym.Longitude)
		if err != nil {
			return nil, fmt.Errorf("error scanning gym row: %w", err)
		}
		gyms = append(gyms, gym)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	if len(gyms) == 0 {
		return []models.Gym{}, nil
	}

	return gyms, nil
}
