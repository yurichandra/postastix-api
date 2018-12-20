package db

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

// Category represents a post category.
type Category struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Posts     []*Post `gorm:"many2many:category_post"`
}

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
	Categories    []*Category `gorm:"many2many:category_post"`
}
