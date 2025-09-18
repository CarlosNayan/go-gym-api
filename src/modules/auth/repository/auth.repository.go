package repository

import (
	"api-gym-on-go/src/config/utils"
	"api-gym-on-go/src/models"
	auth_types "api-gym-on-go/src/modules/auth/types"
	"database/sql"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db}
}

var _ auth_types.IAuthRepository = &AuthRepository{}

func (r *AuthRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	query := `
		SELECT id_user, user_name, email, password, role, created_at
		FROM users
		WHERE email = $1
	`

	row := r.DB.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, utils.WrapError(err)
	}

	return &user, nil
}
