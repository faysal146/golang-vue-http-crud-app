package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// autoCreateTime autoUpdateTime check:age > 13
	ID                string    `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	FirstName         string    `json:"first_name" validate:"required,alpha,max=10,min=1"`
	LastName          string    `json:"last_name" validate:"required,alpha,max=10,min=1"`
	Username          string    `json:"user_name" validate:"required,alphanum,max=10,min=1"`
	Email             string    `json:"email" gorm:"unique" validate:"required,email"`
	Role              string    `json:"role" gorm:"default:user"` //validate:"oneof=user admin"
	Password          string    `json:"password" validate:"required,min=1,max=20"`
	PasswordUpdatedAt time.Time `json:"password_updated_at"`
	Token             string    `json:"token"`
	RefreshToken      string    `json:"refresh_token"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// validate:"eq=user|admin"

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

// temporary
func (u *User) GetUsers(db *gorm.DB) error {
	ud := db.Limit(100).Omit("password", "role", "password_updated_at").Find(&u)
	return ud.Error
}

func (u *User) CreateUser(db *gorm.DB) error {
	return db.Model(&u).Create(u).Error
}
