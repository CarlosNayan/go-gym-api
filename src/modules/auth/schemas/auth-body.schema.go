package auth_schemas

type AuthRequestBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
