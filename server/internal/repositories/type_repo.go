package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type TypeRepository interface {
	CreateType(t *models.Type) error
	UpdateType(t *models.Type) error
	DeleteType(id string) error
	GetAllTypes() ([]models.Type, error)
	GetTypeByID(id string) (*models.Type, error)
}

type typeRepository struct {
	db *gorm.DB
}

func NewTypeRepository(db *gorm.DB) TypeRepository {
	return &typeRepository{db}
}

func (r *typeRepository) CreateType(t *models.Type) error {
	return r.db.Create(t).Error
}

func (r *typeRepository) UpdateType(t *models.Type) error {
	return r.db.Save(t).Error
}

func (r *typeRepository) DeleteType(id string) error {
	return r.db.Delete(&models.Type{}, "id = ?", id).Error
}

func (r *typeRepository) GetAllTypes() ([]models.Type, error) {
	var types []models.Type
	err := r.db.Order("name asc").Find(&types).Error
	return types, err
}

func (r *typeRepository) GetTypeByID(id string) (*models.Type, error) {
	var t models.Type
	err := r.db.First(&t, "id = ?", id).Error
	return &t, err
}
