package repositories

import (
	"errors"
	"server/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	DeleteCategory(id string) error
	GetAllCategories() ([]models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	GetCategoryByID(id string) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) UpdateCategory(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(id string) error {
	return r.db.Delete(&models.Category{}, "id = ?", id).Error
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Order("name asc").Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetCategoryByID(id string) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &category, err
}
