package repository

import "github.com/dewadg/postastix-api/db"

// UserRepository represents repository of User.
type UserRepository struct{}

// Get returns all users.
func (r *UserRepository) Get() []*db.User {
	users := make([]*db.User, 0)

	db.Get().Find(&users)

	return users
}

// Push adds the instance to the repository.
func (r *UserRepository) Push(new *db.User) *db.User {
	db.Get().Create(new)

	return new
}

// Find returns user by ID.
func (r *UserRepository) Find(id uint) *db.User {
	user := new(db.User)
	db.Get().Where("id = ?", id).First(user)

	return user
}

// Delete removes a user by ID.
func (r *UserRepository) Delete(id uint) {
	db.Get().Where("id = ?", id).Delete(db.User{})
}
