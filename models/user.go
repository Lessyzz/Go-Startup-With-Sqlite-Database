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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String() // **UUID'yi string olarak ata**
	return
}
