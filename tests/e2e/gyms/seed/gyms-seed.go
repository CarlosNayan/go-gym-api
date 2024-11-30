package seed

import "api-gym-on-go/models"

func SeedGyms() {
	db := models.SetupDatabase("postgresql://root:admin@127.0.0.1:5432/public?sslmode=disable")
	gym := models.Gym{
		GymName:   "test gym",
		Latitude:  1.23456,
		Longitude: 1.23456,
	}
	db.Create(&gym)
}
