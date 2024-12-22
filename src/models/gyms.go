package models

type Gym struct {
	ID          string  `json:"id_gym"`
	GymName     string  `json:"gym_name"`
	Description *string `json:"description"`
	Phone       *string `json:"phone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
