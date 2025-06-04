package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type LocationRepository interface {
	CreateLocation(location *models.Location) error
	UpdateLocation(location *models.Location) error
	DeleteLocation(id string) error
	GetAllLocations() ([]models.Location, error)
	GetLocationByID(id string) (*models.Location, error)
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db}
}

func (r *locationRepository) CreateLocation(location *models.Location) error {
	return r.db.Create(location).Error
}

func (r *locationRepository) UpdateLocation(location *models.Location) error {
	return r.db.Save(location).Error
}

func (r *locationRepository) DeleteLocation(id string) error {
	return r.db.Delete(&models.Location{}, "id = ?", id).Error
}

func (r *locationRepository) GetAllLocations() ([]models.Location, error) {
	var locations []models.Location
	if err := r.db.Order("name asc").Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func (r *locationRepository) GetLocationByID(id string) (*models.Location, error) {
	var location models.Location
	if err := r.db.First(&location, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &location, nil
}
