package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// autoCreateTime autoUpdateTime check:age > 13
	ID                string    `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Username          string    `json:"user_name"`
	Email             string    `json:"email" gorm:"unique"`
	Role              string    `json:"role" gorm:"default:user"`
	Password          string    `json:"password"`
	PasswordUpdatedAt time.Time `json:"password_updated_at"`
	Token             string    `json:"token"`
	RefreshToken      string    `json:"refresh_token"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
