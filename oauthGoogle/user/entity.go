package user

import (
	"time"
)

// entity for represent table "users" in database
type User struct {
	ID        int    `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string
	Name      string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
