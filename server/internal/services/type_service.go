package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
)

type TypeService interface {
	DeleteType(id string) error
	GetAllTypes() ([]dto.TypeResponse, error)
	CreateType(req dto.CreateTypeRequest) error
	GetTypeByID(id string) (*dto.TypeResponse, error)
	UpdateType(id string, req dto.UpdateTypeRequest) error
}

type typeService struct {
	repo repositories.TypeRepository
}

func NewTypeService(repo repositories.TypeRepository) TypeService {
	return &typeService{repo}
}

func (s *typeService) DeleteType(id string) error {
	_, err := s.repo.GetTypeByID(id)
	if err != nil {
		return customErr.ErrNotFound
	}

	if err := s.repo.DeleteType(id); err != nil {
		return customErr.ErrDeleteFailed
	}

	return nil
}

func (s *typeService) GetAllTypes() ([]dto.TypeResponse, error) {
	types, err := s.repo.GetAllTypes()
	if err != nil {
		return nil, customErr.ErrInternalServer
	}

	var result []dto.TypeResponse
	for _, t := range types {
		result = append(result, dto.TypeResponse{
			ID:   t.ID.String(),
			Name: t.Name,
		})
	}
	return result, nil
}

func (s *typeService) CreateType(req dto.CreateTypeRequest) error {
	t := models.Type{
		Name: req.Name,
	}

	if err := s.repo.CreateType(&t); err != nil {
		return customErr.ErrCreateFailed
	}

	return nil
}

func (s *typeService) GetTypeByID(id string) (*dto.TypeResponse, error) {
	t, err := s.repo.GetTypeByID(id)
	if err != nil {
		return nil, customErr.NewNotFound("Type not found")
	}

	return &dto.TypeResponse{
		ID:   t.ID.String(),
		Name: t.Name,
	}, nil
}

func (s *typeService) UpdateType(id string, req dto.UpdateTypeRequest) error {
	t, err := s.repo.GetTypeByID(id)
	if err != nil {
		return customErr.ErrNotFound
	}

	t.Name = req.Name

	if err := s.repo.UpdateType(t); err != nil {
		return customErr.ErrUpdateFailed
	}
	return nil
}
