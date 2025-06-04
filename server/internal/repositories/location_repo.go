package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type LocationRepository interface {
	DeleteLocation(id string) error
	GetAllLocations() ([]models.Location, error)
	CreateLocation(location *models.Location) error
	UpdateLocation(location *models.Location) error

	GetLocationByID(id string) (*models.Location, error)
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db}
}

func (r *locationRepository) DeleteLocation(id string) error {
	return r.db.Delete(&models.Location{}, "id = ?", id).Error
}

func (r *locationRepository) GetAllLocations() ([]models.Location, error) {
	var locations []models.Location
	err := r.db.Order("name asc").Find(&locations).Error
	return locations, err
}
func (r *locationRepository) CreateLocation(location *models.Location) error {
	return r.db.Create(location).Error
}

func (r *locationRepository) UpdateLocation(location *models.Location) error {
	return r.db.Save(location).Error
}

func (r *locationRepository) GetLocationByID(id string) (*models.Location, error) {
	var location models.Location
	err := r.db.First(&location, "id = ?", id).Error
	return &location, err
}
