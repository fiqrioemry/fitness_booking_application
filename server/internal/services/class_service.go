package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
	"server/pkg/utils"

	"slices"

	"github.com/google/uuid"
)

type ClassService interface {
	DeleteClass(id string) error
	CreateClass(req dto.CreateClassRequest) error
	AddClassGallery(galleries []models.ClassGallery) error
	UpdateClass(id string, req dto.UpdateClassRequest) error
	GetClassByID(id string) (*dto.ClassDetailResponse, error)
	UpdateClassGallery(classID uuid.UUID, keepImages []string, newImageURLs []string) error
	GetAllClasses(params dto.ClassQueryParam) ([]dto.ClassResponse, *dto.PaginationResponse, error)
}

type classService struct {
	repo repositories.ClassRepository
}

func NewClassService(repo repositories.ClassRepository) ClassService {
	return &classService{repo}
}

func (s *classService) CreateClass(req dto.CreateClassRequest) error {
	typeID, err := uuid.Parse(req.TypeID)
	if err != nil {
		return customErr.NewBadRequest("invalid type ID")
	}
	levelID, err := uuid.Parse(req.LevelID)
	if err != nil {
		return customErr.NewBadRequest("invalid level ID")
	}
	locationID, err := uuid.Parse(req.LocationID)
	if err != nil {
		return customErr.NewBadRequest("invalid location ID")
	}
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return customErr.NewBadRequest("invalid category ID")
	}
	subcategoryID, err := uuid.Parse(req.SubcategoryID)
	if err != nil {
		return customErr.NewBadRequest("invalid subcategory ID")
	}

	class := models.Class{
		Title:          req.Title,
		Image:          req.ImageURL,
		Duration:       req.Duration,
		Description:    req.Description,
		AdditionalList: req.Additional,
		TypeID:         typeID,
		LevelID:        levelID,
		LocationID:     locationID,
		CategoryID:     categoryID,
		SubcategoryID:  subcategoryID,
		IsActive:       req.IsActive,
	}

	if err := s.repo.CreateClass(&class); err != nil {
		return customErr.NewInternal("failed to create class", err)
	}

	if len(req.ImageURLs) > 0 {
		var galleries []models.ClassGallery
		for _, url := range req.ImageURLs {
			galleries = append(galleries, models.ClassGallery{ClassID: class.ID, URL: url})
		}
		if err := s.repo.SaveClassGalleries(galleries); err != nil {
			return customErr.NewInternal("failed to save class galleries", err)
		}
	}
	return nil
}

func (s *classService) UpdateClass(id string, req dto.UpdateClassRequest) error {
	class, err := s.repo.GetClassByID(id)
	if err != nil {
		return customErr.NewNotFound("class not found")
	}

	class.Title = req.Title
	class.IsActive = req.IsActive
	class.Duration = req.Duration
	class.Description = req.Description

	if len(req.Additional) > 0 {
		class.AdditionalList = req.Additional
	}

	if req.TypeID != "" {
		typeID, err := uuid.Parse(req.TypeID)
		if err != nil {
			return customErr.NewBadRequest("invalid type ID")
		}
		class.TypeID = typeID
	}
	if req.LevelID != "" {
		levelID, err := uuid.Parse(req.LevelID)
		if err != nil {
			return customErr.NewBadRequest("invalid level ID")
		}
		class.LevelID = levelID
	}
	if req.LocationID != "" {
		locationID, err := uuid.Parse(req.LocationID)
		if err != nil {
			return customErr.NewBadRequest("invalid location ID")
		}
		class.LocationID = locationID
	}
	if req.CategoryID != "" {
		categoryID, err := uuid.Parse(req.CategoryID)
		if err != nil {
			return customErr.NewBadRequest("invalid category ID")
		}
		class.CategoryID = categoryID
	}
	if req.SubcategoryID != "" {
		subcategoryID, err := uuid.Parse(req.SubcategoryID)
		if err != nil {
			return customErr.NewBadRequest("invalid subcategory ID")
		}
		class.SubcategoryID = subcategoryID
	}

	if req.ImageURL != "" {
		_ = utils.DeleteFromCloudinary(class.Image)
		class.Image = req.ImageURL
	}

	if err := s.repo.UpdateClass(class); err != nil {
		return customErr.NewInternal("failed to update class", err)
	}
	return nil
}

func (s *classService) DeleteClass(id string) error {
	class, err := s.repo.GetClassByID(id)
	if err != nil {
		return customErr.NewNotFound("class not found")
	}

	if class.Image != "" {
		_ = utils.DeleteFromCloudinary(class.Image)
	}

	for _, gallery := range class.Galleries {
		if gallery.URL != "" {
			_ = utils.DeleteFromCloudinary(gallery.URL)
		}
	}

	if err := s.repo.DeleteClass(id); err != nil {
		return customErr.NewInternal("failed to delete class", err)
	}
	return nil
}

func (s *classService) GetClassByID(id string) (*dto.ClassDetailResponse, error) {
	class, err := s.repo.GetClassByID(id)
	if err != nil {
		return nil, customErr.NewNotFound("class not found")
	}

	var galleries []string
	for _, g := range class.Galleries {
		galleries = append(galleries, g.URL)
	}

	return &dto.ClassDetailResponse{
		ID:          class.ID.String(),
		Title:       class.Title,
		Image:       class.Image,
		IsActive:    class.IsActive,
		Duration:    class.Duration,
		Description: class.Description,
		Additional:  class.AdditionalList,
		Type:        class.Type.Name,
		Level:       class.Level.Name,
		Location:    class.Location.Name,
		Category:    class.Category.Name,
		Subcategory: class.Subcategory.Name,
		Galleries:   galleries,
		CreatedAt:   class.CreatedAt.Format("2006-01-02"),
	}, nil
}

func (s *classService) GetAllClasses(params dto.ClassQueryParam) ([]dto.ClassResponse, *dto.PaginationResponse, error) {
	classes, total, err := s.repo.GetAllClasses(params)
	if err != nil {
		return nil, nil, customErr.NewNotFound("classes not found")
	}

	var results []dto.ClassResponse
	for _, c := range classes {
		galleries := make([]string, len(c.Galleries))
		for i, g := range c.Galleries {
			galleries[i] = g.URL
		}
		results = append(results, dto.ClassResponse{
			ID:            c.ID.String(),
			Title:         c.Title,
			Image:         c.Image,
			IsActive:      c.IsActive,
			Duration:      c.Duration,
			Galleries:     galleries,
			Description:   c.Description,
			Additional:    c.AdditionalList,
			TypeID:        c.TypeID.String(),
			LevelID:       c.LevelID.String(),
			LocationID:    c.LocationID.String(),
			CategoryID:    c.CategoryID.String(),
			SubcategoryID: c.SubcategoryID.String(),
			CreatedAt:     c.CreatedAt.Format("2006-01-02"),
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return results, pagination, nil
}

func (s *classService) UpdateClassGallery(classID uuid.UUID, keepImages []string, newImageURLs []string) error {
	oldGalleries, err := s.repo.FindGalleriesByClassID(classID)
	if err != nil {
		return customErr.NewNotFound("gallery not found")
	}

	for _, gallery := range oldGalleries {
		if !slices.Contains(keepImages, gallery.URL) {
			_ = utils.DeleteFromCloudinary(gallery.URL)
			_ = s.repo.DeleteClassGalleryByID(gallery.ID.String())
		}
	}

	if len(newImageURLs) > 0 {
		var newGalleries []models.ClassGallery
		for _, url := range newImageURLs {
			newGalleries = append(newGalleries, models.ClassGallery{ClassID: classID, URL: url})
		}
		if err := s.repo.SaveClassGalleries(newGalleries); err != nil {
			return customErr.NewInternal("failed to save new galleries", err)
		}
	}
	return nil
}

func (s *classService) AddClassGallery(galleries []models.ClassGallery) error {
	if err := s.repo.SaveClassGalleries(galleries); err != nil {
		return customErr.NewInternal("failed to save galleries", err)
	}
	return nil
}
