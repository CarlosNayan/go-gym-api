package gyms_schemas

type GymsSearchQuery struct {
	Query string `query:"query" validate:"required"`
}
