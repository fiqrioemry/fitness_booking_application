package repositories

import (
	"errors"
	"server/internal/models"

	"gorm.io/gorm"
)

type SubcategoryRepository interface {
	DeleteSubcategory(id string) error
	GetAllSubcategories() ([]models.Subcategory, error)
	CreateSubcategory(subcategory *models.Subcategory) error
	UpdateSubcategory(subcategory *models.Subcategory) error
	GetSubcategoryByID(id string) (*models.Subcategory, error)
	GetSubcategoriesByCategoryID(categoryID string) ([]models.Subcategory, error)
}

type subcategoryRepository struct {
	db *gorm.DB
}

func NewSubcategoryRepository(db *gorm.DB) SubcategoryRepository {
	return &subcategoryRepository{db}
}

func (r *subcategoryRepository) DeleteSubcategory(id string) error {
	return r.db.Delete(&models.Subcategory{}, "id = ?", id).Error
}

func (r *subcategoryRepository) GetAllSubcategories() ([]models.Subcategory, error) {
	var subcategories []models.Subcategory
	err := r.db.Order("name asc").Find(&subcategories).Error
	return subcategories, err
}

func (r *subcategoryRepository) CreateSubcategory(subcategory *models.Subcategory) error {
	return r.db.Create(subcategory).Error
}

func (r *subcategoryRepository) UpdateSubcategory(subcategory *models.Subcategory) error {
	return r.db.Save(subcategory).Error
}

func (r *subcategoryRepository) GetSubcategoryByID(id string) (*models.Subcategory, error) {
	var subcategory models.Subcategory
	err := r.db.First(&subcategory, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &subcategory, err
}

func (r *subcategoryRepository) GetSubcategoriesByCategoryID(categoryID string) ([]models.Subcategory, error) {
	var subcategories []models.Subcategory
	err := r.db.Where("category_id = ?", categoryID).Find(&subcategories).Error
	return subcategories, err
}
