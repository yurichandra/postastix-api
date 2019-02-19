package repository

import (
	"postastix-api/db"
	"postastix-api/model"

	"github.com/jinzhu/gorm"
)

// CategoryRepository represent struct of category repo
type CategoryRepository struct {
	db *gorm.DB
}

//NewCategoryRepository return new CategoryRepository instance
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		db: db.Get(),
	}
}

// Get return all categories
func (r *CategoryRepository) Get() []model.Category {
	categories := make([]model.Category, 0)

	db.Get().Find(&categories)

	return categories
}

// Find return single category
func (r *CategoryRepository) Find(id uint) model.Category {
	category := new(model.Category)
	db.Get().Where("id = ?", id).First(&category)

	return *category
}

// Push add new category to repository
func (r *CategoryRepository) Push(new *model.Category) {
	db.Get().Create(new)
}

// Delete remove a single category from repository
func (r *CategoryRepository) Delete(id uint) {
	db.Get().Where("id = ?", id).Delete(model.Category{})
}
