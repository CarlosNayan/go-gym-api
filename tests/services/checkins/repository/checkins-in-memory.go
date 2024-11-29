package repository

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/config/utils"
	"api-gym-on-go/src/modules/checkins/interfaces"
	"log"
	"time"
)

type InMemoryCheckinsRepository struct {
	checkins []models.Checkin
	gyms     []models.Gym
}

func NewInMemoryCheckinsRepository() *InMemoryCheckinsRepository {
	return &InMemoryCheckinsRepository{
		checkins: []models.Checkin{
			{
				ID:          "0ebd4f88-d712-4b0f-9278-41d595c690ad",
				IDUser:      "1e2d4f88-d712-4b0f-9278-41d595c690ad",
				IDGym:       "2e2d4f88-d712-4b0f-9278-41d595c690ad",
				CreatedAt:   time.Now(),
				ValidatedAt: nil,
			},
		},
	}
}

var _ interfaces.CheckinsRepository = (*InMemoryCheckinsRepository)(nil)

// CountByUserId implements interfaces.CheckinsRepository.
func (i *InMemoryCheckinsRepository) CountByUserId(id_user string) (int64, error) {
	var count int64

	for _, item := range i.checkins {
		if item.IDUser == id_user {
			count++
		}
	}

	return count, nil
}

// CreateCheckin implements interfaces.CheckinsRepository.
func (i *InMemoryCheckinsRepository) CreateCheckin(checkin *models.Checkin) error {
	i.checkins = append(i.checkins, *checkin)
	return nil
}

// FindCheckinById implements interfaces.CheckinsRepository.
func (i *InMemoryCheckinsRepository) FindCheckinById(id_checkin string) (*models.Checkin, error) {
	var checkin models.Checkin

	for _, item := range i.checkins {
		if item.ID == id_checkin {
			checkin = item
			break
		}
	}

	return &checkin, nil
}

func (i *InMemoryCheckinsRepository) FindCheckinByIdOnDate(id_user string) (*models.Checkin, error) {
	var checkin models.Checkin

	now, err := utils.NewMoment()
	if err != nil {
		log.Fatalf("Erro ao criar o data: %v", err)
	}

	startOfDay := now.StartOf("day").ToDate()
	endOfDay := now.EndOf("day").ToDate()

	for _, item := range i.checkins {
		if item.IDUser == id_user &&
			!item.CreatedAt.Before(startOfDay) &&
			!item.CreatedAt.After(endOfDay) {
			checkin = item
			break
		}
	}

	return &checkin, nil
}

func (i *InMemoryCheckinsRepository) FindGymByID(id_gym string) (*models.Gym, error) {
	var gym models.Gym

	for _, item := range i.gyms {
		if item.ID == id_gym {
			gym = item
			break
		}
	}

	return &gym, nil
}

func (i *InMemoryCheckinsRepository) ListAllCheckinsHistoryOfUser(id_user string, page int) ([]models.Checkin, error) {
	var checkins []models.Checkin

	for _, item := range i.checkins {
		if item.IDUser == id_user {
			checkins = append(checkins, item)
		}
	}

	return checkins, nil
}

func (i *InMemoryCheckinsRepository) UpdateCheckin(id_checkin string) (*models.Checkin, error) {
	var updatedCheckin models.Checkin
	now := time.Now()

	for _, item := range i.checkins {
		if item.ID == id_checkin {
			item.ValidatedAt = &now
			updatedCheckin = item
			break
		}
	}

	return &updatedCheckin, nil
}
