package repository

import (
	"postastix-api/db"
	"postastix-api/model"

	"github.com/jinzhu/gorm"
)

// UserRepository represents repository of User.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns new user repository instance.
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.Get(),
	}
}

// Get returns all users.
func (r *UserRepository) Get() []model.User {
	users := make([]model.User, 0)

	db.Get().Find(&users)

	return users
}

// Push adds the instance to the repository.
func (r *UserRepository) Push(new *model.User) {
	db.Get().Create(new)
}

// Find returns user by ID.
func (r *UserRepository) Find(id uint) model.User {
	user := new(model.User)
	db.Get().Where("id = ?", id).First(&user)

	return *user
}

// Delete removes a user by ID.
func (r *UserRepository) Delete(id uint) {
	db.Get().Where("id = ?", id).Delete(model.User{})
}
