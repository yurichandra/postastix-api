package repository

import (
	"postastix-api/model"
)

// UserRepositoryContract represents contract for UserRepository
type UserRepositoryContract interface {
	Get() []model.User
	Push(new *model.User)
	Find(id uint) model.User
	Delete(id uint)
}

// CategoryRepositoryContract represents contract for CategoryRepository
type CategoryRepositoryContract interface {
	Get() []model.Category
	Find(id uint) model.Category
	Push(new *model.Category)
	Delete(id uint)
}
