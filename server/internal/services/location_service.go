package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"

	"github.com/google/uuid"
)

type LocationService interface {
	DeleteLocation(id string) error
	GetAllLocations() ([]dto.LocationResponse, error)
	CreateLocation(req dto.CreateLocationRequest) error
	GetLocationByID(id string) (*dto.LocationResponse, error)
	UpdateLocation(id string, req dto.UpdateLocationRequest) error
}

type locationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) LocationService {
	return &locationService{repo}
}

func (s *locationService) CreateLocation(req dto.CreateLocationRequest) error {
	location := models.Location{
		ID:          uuid.New(),
		Name:        req.Name,
		Address:     req.Address,
		GeoLocation: req.GeoLocation,
	}

	if err := s.repo.CreateLocation(&location); err != nil {
		return customErr.ErrCreateFailed
	}
	return nil
}

func (s *locationService) UpdateLocation(id string, req dto.UpdateLocationRequest) error {
	location, err := s.repo.GetLocationByID(id)
	if err != nil {
		return customErr.ErrNotFound
	}

	location.Name = req.Name
	location.Address = req.Address
	location.GeoLocation = req.GeoLocation

	if err := s.repo.UpdateLocation(location); err != nil {
		return customErr.ErrUpdateFailed
	}
	return nil
}

func (s *locationService) DeleteLocation(id string) error {
	_, err := s.repo.GetLocationByID(id)
	if err != nil {
		return customErr.ErrNotFound
	}

	if err := s.repo.DeleteLocation(id); err != nil {
		return customErr.ErrDeleteFailed
	}
	return nil
}

func (s *locationService) GetAllLocations() ([]dto.LocationResponse, error) {
	locations, err := s.repo.GetAllLocations()
	if err != nil {
		return nil, customErr.NewInternal("Failed to fetch locations", err)
	}

	var result []dto.LocationResponse
	for _, l := range locations {
		result = append(result, dto.LocationResponse{
			ID:          l.ID.String(),
			Name:        l.Name,
			Address:     l.Address,
			GeoLocation: l.GeoLocation,
		})
	}
	return result, nil
}

func (s *locationService) GetLocationByID(id string) (*dto.LocationResponse, error) {
	location, err := s.repo.GetLocationByID(id)
	if err != nil {
		return nil, customErr.NewNotFound("Location not found")
	}

	return &dto.LocationResponse{
		ID:          location.ID.String(),
		Name:        location.Name,
		Address:     location.Address,
		GeoLocation: location.GeoLocation,
	}, nil
}
