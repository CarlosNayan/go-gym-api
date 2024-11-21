package schema

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
	ID           string    `gorm:"column:id_user;primaryKey;type:uuid;default:gen_random_uuid()"`
	UserName     string    `gorm:"column:user_name;type:varchar(255);not null"`
	Email        string    `gorm:"column:email;uniqueIndex;type:varchar(255);not null"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null"`
	Role         Role      `gorm:"column:role;type:varchar(10);default:'MEMBER';not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:now()"`
	Checkins     []Checkin `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE"`
}

type Gym struct {
	ID          string    `gorm:"column:id_gym;primaryKey;type:uuid;default:gen_random_uuid()"`
	GymName     string    `gorm:"column:gym_name;type:varchar(255);not null"`
	Description *string   `gorm:"column:description;type:text"`
	Phone       *string   `gorm:"column:phone;type:varchar(15)"`
	Latitude    float64   `gorm:"column:latitude;type:decimal(10,8);not null"`
	Longitude   float64   `gorm:"column:longitude;type:decimal(11,8);not null"`
	Checkins    []Checkin `gorm:"foreignKey:IDGym;constraint:OnDelete:CASCADE"`
}

type Checkin struct {
	ID          string     `gorm:"column:id_checkin;primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp;default:now()"`
	ValidatedAt *time.Time `gorm:"column:validated_at;type:timestamp"`
	IDUser      string     `gorm:"column:id_user;type:uuid;not null"`
	IDGym       string     `gorm:"column:id_gym;type:uuid;not null"`

	User User `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE;referencedKey:id_user"`
	Gym  Gym  `gorm:"foreignKey:IDGym;constraint:OnDelete:CASCADE;referencedKey:id_gym"`
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
