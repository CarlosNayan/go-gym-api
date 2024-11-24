package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin  Role = "ADMIN"
	RoleMember Role = "MEMBER"
)

type User struct {
	ID           string    `gorm:"column:id_user;primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserName     string    `gorm:"column:user_name;type:varchar(255);not null" json:"user_name"`
	Email        string    `gorm:"column:email;uniqueIndex;type:varchar(255);not null" json:"email"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null" json:"password_hash"`
	Role         Role      `gorm:"column:role;type:varchar(10);default:'MEMBER';not null" json:"role"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:now()" json:"createdAt"`
	Checkins     []Checkin `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE" json:"checkins"`
}

type Gym struct {
	ID          string  `gorm:"column:id_gym;primaryKey;type:uuid;default:gen_random_uuid()" json:"id_gym"`
	GymName     string  `gorm:"column:gym_name;type:varchar(255);not null" json:"gym_name"`
	Description *string `gorm:"column:description;type:text" json:"description"`
	Phone       *string `gorm:"column:phone;type:varchar(15)" json:"phone"`
	Latitude    float64 `gorm:"column:latitude;type:decimal(10,8);not null" json:"latitude"`
	Longitude   float64 `gorm:"column:longitude;type:decimal(11,8);not null" json:"longitude"`

	Checkins []Checkin `gorm:"foreignKey:IDGym;constraint:OnDelete:CASCADE"`
}

type Checkin struct {
	ID          string     `gorm:"column:id_checkin;primaryKey;type:uuid;default:gen_random_uuid()" json:"id_checkin"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp;default:now()" json:"created_at"`
	ValidatedAt *time.Time `gorm:"column:validated_at;type:timestamp" json:"validated_at"`
	IDUser      string     `gorm:"column:id_user;type:uuid;not null" json:"id_user"`
	IDGym       string     `gorm:"column:id_gym;type:uuid;not null" json:"id_gym"`

	User User `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE;referencedKey:id_user" json:"-"`
	Gym  Gym  `gorm:"foreignKey:IDGym;constraint:OnDelete:CASCADE;referencedKey:id_gym" json:"-"`
}

func SetupDatabase(database string) *gorm.DB {
	dsn := database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&User{}, &Gym{}, &Checkin{})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return db
}
