package model

import "time"

// Post represents a post.
type Post struct {
	ID            uint `gorm:"primary_key"`
	Title         string
	HTMLContent   string
	MDContent     string
	PublishedAt   time.Time
	RepublishedAt *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	UserID        uint
	Categories    []Category `gorm:"many2many:category_post"`
}
