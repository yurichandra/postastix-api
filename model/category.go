package model

import "time"

// Category represents a post category.
type Category struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Posts     []Post `gorm:"many2many:category_post"`
}
