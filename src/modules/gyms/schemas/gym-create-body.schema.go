package gyms_schemas

type GymCreateBody struct {
	GymName     string  `json:"gym_name" validate:"required"`
	Description *string `json:"description"`
	Phone       *string `json:"phone"`
	Latitude    float64 `json:"latitude" validate:"required,latitude"`
	Longitude   float64 `json:"longitude" validate:"required,longitude"`
}
