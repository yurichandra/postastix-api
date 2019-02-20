package service

import (
	"errors"
	"postastix-api/db"
	"postastix-api/model"
	"postastix-api/repository"
)

// CategoryService represents service layer for category
type CategoryService struct {
	repo repository.CategoryRepositoryContract
}

// NewCategoryService create service for a category
func NewCategoryService() *CategoryService {
	return &CategoryService{
		repo: repository.NewCategoryRepository(),
	}
}

// Get return all categories
func (r *CategoryService) Get() []model.Category {
	return r.repo.Get()
}

// Find return single category
func (r *CategoryService) Find(id uint) (model.Category, error) {
	category := r.repo.Find(id)

	if category.ID == 0 {
		return model.Category{}, errors.New("Category not found")
	}

	return category, nil
}

func (r *CategoryService) isNameUnique(name string) bool {
	category := model.Category{}

	db.Get().Where("name = ?", name).First(&category)

	return category.ID == 0
}

// Create is save new category to DB
func (r *CategoryService) Create(name string) (model.Category, error) {
	if !r.isNameUnique(name) {
		return model.Category{}, errors.New("Category already available")
	}

	newCategory := model.Category{
		Name: name,
	}

	r.repo.Push(&newCategory)

	return newCategory, nil
}

// Update is update new category
func (r *CategoryService) Update(id uint, name string) (model.Category, error) {
	category := r.repo.Find(id)

	if category.ID == 0 {
		return model.Category{}, errors.New("Category not found")
	}

	if !r.isNameUnique(name) {
		return model.Category{}, errors.New("Category already available")
	}

	category.Name = name
	db.Get().Save(&category)

	return category, nil
}

// Delete is delete category in db
func (r *CategoryService) Delete(id uint) error {
	category := r.repo.Find(id)

	if category.ID == 0 {
		return errors.New("Category not found")
	}

	r.repo.Delete(id)

	return nil
}
