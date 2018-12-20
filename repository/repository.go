package repository

import "github.com/dewadg/postastix-api/db"

// UserRepositoryContract represents contract for UserRepository
type UserRepositoryContract interface {
	Get() []*db.User
	Push(new *db.User) *db.User
	Find(id uint) *db.User
	Delete(id uint)
}

// GetUserRepository returns new userRepository instance.
func GetUserRepository() UserRepositoryContract {
	return new(userRepository)
}
