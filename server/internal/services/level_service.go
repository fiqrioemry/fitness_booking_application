package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
)

type LevelService interface {
	DeleteLevel(id string) error
	GetAllLevels() ([]dto.LevelResponse, error)
	CreateLevel(req dto.CreateLevelRequest) error
	GetLevelByID(id string) (*dto.LevelResponse, error)
	UpdateLevel(id string, req dto.UpdateLevelRequest) error
}

type levelService struct {
	repo repositories.LevelRepository
}

func NewLevelService(repo repositories.LevelRepository) LevelService {
	return &levelService{repo}
}

func (s *levelService) DeleteLevel(id string) error {
	_, err := s.repo.GetLevelByID(id)
	if err != nil {
		return customErr.NewNotFound("level not found")
	}

	if err := s.repo.DeleteLevel(id); err != nil {
		return customErr.NewInternal("failed to delete level", err)
	}

	return nil
}

func (s *levelService) GetAllLevels() ([]dto.LevelResponse, error) {
	levels, err := s.repo.GetAllLevels()
	if err != nil {
		return nil, customErr.NewInternal("failed to fetch levels", err)
	}

	var result []dto.LevelResponse
	for _, l := range levels {
		result = append(result, dto.LevelResponse{
			ID:   l.ID.String(),
			Name: l.Name,
		})
	}
	return result, nil
}

func (s *levelService) CreateLevel(req dto.CreateLevelRequest) error {
	level := models.Level{
		Name: req.Name,
	}

	if err := s.repo.CreateLevel(&level); err != nil {
		return customErr.NewInternal("failed to create level", err)
	}

	return nil
}

func (s *levelService) GetLevelByID(id string) (*dto.LevelResponse, error) {
	level, err := s.repo.GetLevelByID(id)
	if err != nil {
		return nil, customErr.NewNotFound("level not found")
	}

	return &dto.LevelResponse{
		ID:   level.ID.String(),
		Name: level.Name,
	}, nil
}

func (s *levelService) UpdateLevel(id string, req dto.UpdateLevelRequest) error {
	level, err := s.repo.GetLevelByID(id)
	if err != nil {
		return customErr.NewNotFound("level not found")
	}

	level.Name = req.Name

	if err := s.repo.UpdateLevel(level); err != nil {
		return customErr.NewInternal("failed to update level", err)
	}

	return nil
}
