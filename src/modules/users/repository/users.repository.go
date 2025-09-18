package repository

import (
	"database/sql"

	"api-gym-on-go/src/config/utils"
	"api-gym-on-go/src/models"
	users_schemas "api-gym-on-go/src/modules/users/schemas"
	users_types "api-gym-on-go/src/modules/users/types"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

var _ users_types.UserRepository = &UserRepository{}

func (r *UserRepository) GetProfileById(id string) (*models.User, error) {
	var user models.User

	query := `
		SELECT id_user, user_name, email, role, created_at
		FROM users
		WHERE id_user = $1
		LIMIT 1
	`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, utils.WrapError(err)
	}

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Role, &user.CreatedAt)
		if err != nil {
			return nil, utils.WrapError(err)
		}
	}

	return &user, nil
}

func (r *UserRepository) UserEmailVerify(email string) (*string, error) {
	var user models.User

	query := `
		SELECT email
		FROM users
		WHERE email = $1
	`

	row := r.DB.QueryRow(query, email)
	err := row.Scan(&user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, utils.WrapError(err)
	}

	return &user.Email, nil
}

func (r *UserRepository) CreateUser(user *users_schemas.UserCreateBody) (*models.User, error) {
	var createdUser models.User

	id := uuid.New()

	query := `
		INSERT INTO users (id_user, user_name, email, password, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id_user, user_name, email, role
	`

	err := r.DB.QueryRow(query, id, user.UserName, user.Email, user.Password, user.Role).Scan(&createdUser.ID, &createdUser.UserName, &createdUser.Email, &createdUser.Role)
	if err != nil {
		return nil, utils.WrapError(err)
	}

	return &createdUser, nil
}
