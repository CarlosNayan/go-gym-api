package seed

import (
	"api-gym-on-go/models"
)

func SeedCheckins() {
	db := models.SetupDatabase("postgresql://root:admin@127.0.0.1:5432/public?sslmode=disable")
	gym := models.Gym{
		ID:        "2e2d4f88-d712-4b0f-9278-41d595c690ad",
		GymName:   "test gym",
		Latitude:  1.23456,
		Longitude: 1.23456,
	}
	checkin := models.Checkin{
		ID:     "0ebd4f88-d712-4b0f-9278-41d595c690ad",
		IDUser: "1e2d4f88-d712-4b0f-9278-41d595c690ad",
		IDGym:  "2e2d4f88-d712-4b0f-9278-41d595c690ad",
	}
	err := db.Create(&gym).Error
	if err != nil {
		panic(err)
	}
	err = db.Create(&checkin).Error
	if err != nil {
		panic(err)
	}
}
