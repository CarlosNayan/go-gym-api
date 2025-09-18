package users_schemas

type UserCreateBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
	Role     string `json:"role" validate:"required"`
}
