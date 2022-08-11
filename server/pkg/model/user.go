package model

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	// autoCreateTime autoUpdateTime check:age > 13
	ID                string    `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Username          string    `json:"user_name" gorm:"unique"`
	Email             string    `json:"email" gorm:"unique"`
	Role              string    `json:"-" gorm:"default:user"` //validate:"oneof=user admin"
	Password          string    `json:"-"`
	PasswordUpdatedAt time.Time `json:"-"`
	Token             string    `json:"token"`
	RefreshToken      string    `json:"refresh_token"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// incoming user data
type LoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}
type RegisterBody struct {
	FirstName string `json:"first_name" validate:"required,alpha,max=10,min=1"`
	LastName  string `json:"last_name" validate:"required,alpha,max=10,min=1"`
	Username  string `json:"user_name" gorm:"unique" validate:"required,alphanum,max=10,min=3"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=20"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// make email address small letter
	u.Email = strings.ToLower(u.Email)
	u.Password = u.HashPassword(u.Password)
	// hash password
	return
}

func (u *User) HashPassword(p string) string {
	passwordByte, _ := bcrypt.GenerateFromPassword([]byte(p), 8)
	return string(passwordByte)
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
