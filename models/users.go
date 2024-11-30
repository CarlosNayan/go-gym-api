package models

import "time"

type Role string

const (
	RoleAdmin  Role = "ADMIN"
	RoleMember Role = "MEMBER"
)

type User struct {
	ID           string    `json:"id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Role         Role      `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}
