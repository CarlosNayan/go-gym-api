package repository

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/utils"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
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
