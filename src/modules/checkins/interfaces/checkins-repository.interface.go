package interfaces

import "api-gym-on-go/src/models"

type CheckinsRepository interface {
	CreateCheckin(checkin *models.Checkin) error
	FindCheckinByIdOnDate(id_user string) (*models.Checkin, error)
	FindCheckinById(id_checkin string) (*models.Checkin, error)
	UpdateCheckin(id_checkin string) (*models.Checkin, error)
	CountByUserId(id_user string) (int64, error)
	ListAllCheckinsHistoryOfUser(id_user string, page int) ([]models.Checkin, error)
	FindGymByID(id_gym string) (*models.Gym, error)
}
