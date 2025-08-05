package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint
	Name     *string `gorm:"not null"`
	Email    string  `gorm:"unique"`
	Password string
	Username string `gorm:"unique;not null"`
}
