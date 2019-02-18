package repository

import (
	"postastix-api/model"
)

// UserRepositoryContract represents contract for UserRepository
type UserRepositoryContract interface {
	Get() []model.User
	Push(new model.User)
	Find(id uint) model.User
	Delete(id uint)
}
