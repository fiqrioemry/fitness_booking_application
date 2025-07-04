package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"

	"github.com/google/uuid"
)

type SubcategoryService interface {
	DeleteSubcategory(id string) error
	GetAllSubcategories() ([]dto.SubcategoryResponse, error)
	CreateSubcategory(req dto.CreateSubcategoryRequest) error
	GetSubcategoryByID(id string) (*dto.SubcategoryResponse, error)
	UpdateSubcategory(id string, req dto.UpdateSubcategoryRequest) error
	GetSubcategoriesByCategoryID(categoryID string) ([]dto.SubcategoryResponse, error)
}

type subcategoryService struct {
	repo repositories.SubcategoryRepository
}

func NewSubcategoryService(repo repositories.SubcategoryRepository) SubcategoryService {
	return &subcategoryService{repo}
}

func (s *subcategoryService) CreateSubcategory(req dto.CreateSubcategoryRequest) error {
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return customErr.NewBadRequest("invalid category ID")
	}

	subcategory := models.Subcategory{
		ID:         uuid.New(),
		Name:       req.Name,
		CategoryID: categoryID,
	}

	if err := s.repo.CreateSubcategory(&subcategory); err != nil {
		return customErr.NewInternal("failed to create subcategory", err)
	}

	return nil
}

func (s *subcategoryService) UpdateSubcategory(id string, req dto.UpdateSubcategoryRequest) error {
	subcategory, err := s.repo.GetSubcategoryByID(id)
	if err != nil {
		return customErr.NewNotFound("subcategory not found")
	}

	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return customErr.NewBadRequest("invalid category ID")
	}

	subcategory.Name = req.Name
	subcategory.CategoryID = categoryID

	if err := s.repo.UpdateSubcategory(subcategory); err != nil {
		return customErr.NewInternal("failed to update subcategory", err)
	}

	return nil
}

func (s *subcategoryService) DeleteSubcategory(id string) error {
	_, err := s.repo.GetSubcategoryByID(id)
	if err != nil {
		return customErr.NewNotFound("subcategory not found")
	}

	if err := s.repo.DeleteSubcategory(id); err != nil {
		return customErr.NewInternal("failed to delete subcategory", err)
	}

	return nil
}

func (s *subcategoryService) GetAllSubcategories() ([]dto.SubcategoryResponse, error) {
	subcategories, err := s.repo.GetAllSubcategories()
	if err != nil {
		return nil, customErr.NewInternal("failed to fetch subcategories", err)
	}

	var result []dto.SubcategoryResponse
	for _, s := range subcategories {
		result = append(result, dto.SubcategoryResponse{
			ID:         s.ID.String(),
			Name:       s.Name,
			CategoryID: s.CategoryID.String(),
		})
	}

	return result, nil
}

func (s *subcategoryService) GetSubcategoryByID(id string) (*dto.SubcategoryResponse, error) {
	subcategory, err := s.repo.GetSubcategoryByID(id)
	if err != nil {
		return nil, customErr.NewNotFound("subcategory not found")
	}

	return &dto.SubcategoryResponse{
		ID:         subcategory.ID.String(),
		Name:       subcategory.Name,
		CategoryID: subcategory.CategoryID.String(),
	}, nil
}

func (s *subcategoryService) GetSubcategoriesByCategoryID(categoryID string) ([]dto.SubcategoryResponse, error) {
	subcategories, err := s.repo.GetSubcategoriesByCategoryID(categoryID)
	if err != nil {
		return nil, customErr.NewInternal("failed to fetch subcategories by category ID", err)
	}

	var result []dto.SubcategoryResponse
	for _, s := range subcategories {
		result = append(result, dto.SubcategoryResponse{
			ID:         s.ID.String(),
			Name:       s.Name,
			CategoryID: s.CategoryID.String(),
		})
	}

	return result, nil
}
