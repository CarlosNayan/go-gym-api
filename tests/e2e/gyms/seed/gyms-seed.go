package seed

import (
	"api-gym-on-go/src/config/database"
	"api-gym-on-go/src/models"
)

func SeedGyms() {
	db := database.SetupDatabase("postgresql://root:admin@127.0.0.1:5432/public?sslmode=disable")
	gym := models.Gym{
		ID:        "2e2d4f88-d712-4b0f-9278-41d595c690ad",
		GymName:   "test gym",
		Latitude:  1.23456,
		Longitude: 1.23456,
	}

	query := `
		INSERT INTO gyms
		(id_gym, gym_name, latitude, longitude)
		VALUES
		($1, $2, $3, $4)
		RETURNING id_gym
	`

	err := db.QueryRow(query, gym.ID, gym.GymName, gym.Latitude, gym.Longitude).Scan(&gym.ID)
	if err != nil {
		panic(err)
	}
}
