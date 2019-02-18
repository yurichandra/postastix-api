package model

import "time"

// User represents a user.
type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	FullName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
