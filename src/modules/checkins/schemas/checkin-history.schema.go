package checkin_schemas

type CheckinValidateQuery struct {
	Page int `query:"page" validate:"required,min=1"`
}
