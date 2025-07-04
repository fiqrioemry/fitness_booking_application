package repositories

import (
	"errors"
	"server/internal/models"

	"gorm.io/gorm"
)

type LevelRepository interface {
	DeleteLevel(id string) error
	CreateLevel(l *models.Level) error
	UpdateLevel(l *models.Level) error
	GetAllLevels() ([]models.Level, error)
	GetLevelByID(id string) (*models.Level, error)
}

type levelRepository struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) LevelRepository {
	return &levelRepository{db}
}

func (r *levelRepository) DeleteLevel(id string) error {
	return r.db.Delete(&models.Level{}, "id = ?", id).Error
}

func (r *levelRepository) CreateLevel(l *models.Level) error {
	return r.db.Create(l).Error
}

func (r *levelRepository) UpdateLevel(l *models.Level) error {
	return r.db.Save(l).Error
}

func (r *levelRepository) GetAllLevels() ([]models.Level, error) {
	var levels []models.Level
	err := r.db.Order("name asc").Find(&levels).Error
	return levels, err
}

func (r *levelRepository) GetLevelByID(id string) (*models.Level, error) {
	var level models.Level
	err := r.db.First(&level, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &level, err

}
