package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint
	Name      string
	Username  *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
