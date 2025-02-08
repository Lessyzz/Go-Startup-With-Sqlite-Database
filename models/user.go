package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"unique"`
	Password string
}

type UserRegisterDTO struct {
	Username string
	Email    string
	Password string
}

type UserLoginDTO struct {
	Username string
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String() // UUID
	return
}
