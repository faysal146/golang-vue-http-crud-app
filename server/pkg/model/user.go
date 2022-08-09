package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	// autoCreateTime autoUpdateTime check:age > 13
	ID                string    `json:"id" gorm:"primaryKey;default:uuid_generate_v3(); not null"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Username          string    `json:"user_name"`
	Email             string    `json:"email" gorm:"unique"`
	Role              string    `json:"role" gorm:"<-:false;default:user"`
	Password          string    `json:"password" gorm:"->:false"`
	PasswordUpdatedAt time.Time `json:"password_updated_at"`
	Token             string    `json:"token"`
	RefreshToken      string    `json:"refresh_token"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Email == "admin@example.com" {
		u.Role = "admin"
	}
	return
}
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("user created. user id is :", u.ID)
	return
}
