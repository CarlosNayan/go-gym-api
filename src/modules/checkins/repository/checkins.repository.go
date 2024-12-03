package repository

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/utils"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type CheckinRepository struct {
	DB *sql.DB
}

func NewCheckinRepository(db *sql.DB) *CheckinRepository {
	return &CheckinRepository{DB: db}
}

func (cr *CheckinRepository) CreateCheckin(checkin *models.Checkin) error {
	id := uuid.New()

	query := `
		INSERT INTO checkins
		(id_checkin, id_user, id_gym)
		VALUES
		($1, $2, $3)
	`

	_, err := cr.DB.Exec(query, id, checkin.IDUser, checkin.IDGym)
	if err != nil {
		log.Println(utils.WrapError(err))
		return utils.WrapError(err)
	}

	return nil
}

func (cr *CheckinRepository) FindCheckinByIdOnDate(id_user string) (*models.Checkin, error) {
	var checkin models.Checkin

	now, err := utils.NewMoment()
	if err != nil {
		utils.WrapError(err)
	}

	startOfDay := now.StartOf("day").Format()
	endOfDay := now.EndOf("day").Format()

	query := `
		SELECT id_checkin, id_user, id_gym, created_at 
		FROM checkins 
		WHERE id_user = $1 
		AND created_at BETWEEN $2 AND $3
	`

	row := cr.DB.QueryRow(query, id_user, startOfDay, endOfDay)
	err = row.Scan(&checkin.ID, &checkin.IDUser, &checkin.IDGym, &checkin.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, utils.WrapError(err)
	}

	return &checkin, nil
}

func (cr *CheckinRepository) FindCheckinById(id_checkin string) (*models.Checkin, error) {
	var checkin models.Checkin

	query := `
		SELECT id_checkin, id_user, id_gym, created_at, validated_at
		FROM checkins 
		WHERE id_checkin = $1
	`

	row := cr.DB.QueryRow(query, id_checkin)
	err := row.Scan(&checkin.ID, &checkin.IDUser, &checkin.IDGym, &checkin.CreatedAt, &checkin.ValidatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, utils.WrapError(err)
	}

	return &checkin, nil
}

func (cr *CheckinRepository) UpdateCheckin(id_checkin string) (*models.Checkin, error) {
	var updatedCheckin models.Checkin

	query := `
		UPDATE checkins
		SET validated_at = now()
		WHERE id_checkin = $1
		RETURNING id_checkin, id_user, id_gym, validated_at, created_at
	`

	row := cr.DB.QueryRow(query, id_checkin)
	if row.Err() != nil {
		return nil, utils.WrapError(row.Err())
	}

	err := row.Scan(&updatedCheckin.ID, &updatedCheckin.IDUser, &updatedCheckin.IDGym, &updatedCheckin.CreatedAt, &updatedCheckin.ValidatedAt)
	if err != nil {
		return nil, utils.WrapError(err)
	}

	return &updatedCheckin, nil
}

func (cr *CheckinRepository) CountByUserId(id_user string) (int64, error) {
	var count int64

	query := `
		SELECT COUNT(*) FROM checkins 
		WHERE id_user = $1
	`

	row := cr.DB.QueryRow(query, id_user)
	err := row.Scan(&count)
	if err != nil {
		return 0, utils.WrapError(err)
	}

	return count, err
}

func (cr *CheckinRepository) ListAllCheckinsHistoryOfUser(id_user string, page int) ([]models.Checkin, error) {
	var checkins []models.Checkin

	query := `
		SELECT id_checkin, id_user, id_gym, created_at 
		FROM checkins 
		WHERE id_user = $1
		LIMIT 10
		OFFSET $2
	`

	rows, err := cr.DB.Query(query, id_user, (page-1)*10)
	if err != nil {
		return nil, utils.WrapError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var checkin models.Checkin
		err = rows.Scan(&checkin.ID, &checkin.IDUser, &checkin.IDGym, &checkin.CreatedAt)
		if err != nil {
			return nil, utils.WrapError(err)
		}
		checkins = append(checkins, checkin)
	}

	if err = rows.Err(); err != nil {
		return nil, utils.WrapError(err)
	}

	return checkins, err
}

func (cr *CheckinRepository) FindGymByID(id_gym string) (*models.Gym, error) {
	var gym models.Gym

	query := `
		SELECT id_gym, gym_name, description, latitude, longitude FROM gyms
		WHERE id_gym = $1
	`

	row := cr.DB.QueryRow(query, id_gym)
	err := row.Scan(&gym.ID, &gym.GymName, &gym.Description, &gym.Latitude, &gym.Longitude)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, utils.WrapError(err)
	}

	return &gym, nil
}
