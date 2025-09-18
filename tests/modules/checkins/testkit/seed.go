package checkins_e2e_test_kit

import (
	"api-gym-on-go/src/models"
	"api-gym-on-go/tests/utils"
)

func SeedCheckins(preCreateCheckin bool) {
	db := utils.SetupDatabase("postgresql://root:admin@127.0.0.1:5432/public?sslmode=disable")
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
	`

	err := db.QueryRow(query, gym.ID, gym.GymName, gym.Latitude, gym.Longitude)
	if err.Err() != nil {
		panic(err)
	}
	if preCreateCheckin {
		checkin := models.Checkin{
			ID:     "0ebd4f88-d712-4b0f-9278-41d595c690ad",
			IDUser: "1e2d4f88-d712-4b0f-9278-41d595c690ad",
			IDGym:  "2e2d4f88-d712-4b0f-9278-41d595c690ad",
		}

		query = `
			INSERT INTO checkins
			(id_checkin, id_user, id_gym, created_at)
			VALUES
			($1, $2, $3, NOW())
		`
		err := db.QueryRow(query, checkin.ID, checkin.IDUser, checkin.IDGym)
		if err.Err() != nil {
			panic(err)
		}
	}
}
