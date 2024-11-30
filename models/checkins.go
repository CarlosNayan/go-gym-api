package models

import "time"

type Checkin struct {
	ID          string     `json:"id_checkin"`
	CreatedAt   time.Time  `json:"created_at"`
	ValidatedAt *time.Time `json:"validated_at"`
	IDUser      string     `json:"id_user"`
	IDGym       string     `json:"id_gym"`
}
