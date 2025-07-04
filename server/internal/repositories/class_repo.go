package repositories

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassRepository interface {
	DeleteClass(id string) error
	UpdateClass(class *models.Class) error
	CreateClass(class *models.Class) error
	DeleteClassGalleryByID(id string) error
	GetClassByID(id string) (*models.Class, error)
	SaveClassGalleries(galleries []models.ClassGallery) error
	GetAllClasses(params dto.ClassQueryParam) ([]models.Class, int64, error)
	FindGalleriesByClassID(classID uuid.UUID) ([]models.ClassGallery, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db}
}

func (r *classRepository) DeleteClass(id string) error {
	return r.db.Delete(&models.Class{}, "id = ?", id).Error
}

func (r *classRepository) CreateClass(class *models.Class) error {
	return r.db.Create(class).Error
}

func (r *classRepository) UpdateClass(class *models.Class) error {
	return r.db.Save(class).Error
}
func (r *classRepository) DeleteClassGalleryByID(id string) error {
	return r.db.Delete(&models.ClassGallery{}, "id = ?", id).Error
}

func (r *classRepository) GetClassByID(id string) (*models.Class, error) {
	var class models.Class
	err := r.db.
		Preload("Type").
		Preload("Level").
		Preload("Category").
		Preload("Subcategory").
		Preload("Location").
		Preload("Galleries").
		First(&class, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &class, err
}

func (r *classRepository) SaveClassGalleries(galleries []models.ClassGallery) error {
	return r.db.Create(&galleries).Error
}

func (r *classRepository) GetAllClasses(params dto.ClassQueryParam) ([]models.Class, int64, error) {
	var classes []models.Class
	var count int64

	db := r.db.Model(&models.Class{}).Preload("Galleries")

	// search
	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("title LIKE ? OR description LIKE ?", like, like)
	}

	// status
	if params.Status != "" && params.Status != "all" {
		if params.Status == "active" {
			db = db.Where("is_active = ?", true)
		} else if params.Status == "inactive" {
			db = db.Where("is_active = ?", false)
		}
	}

	// filters
	if params.TypeID != "" {
		db = db.Where("type_id = ?", params.TypeID)
	}
	if params.CategoryID != "" {
		db = db.Where("category_id = ?", params.CategoryID)
	}
	if params.LevelID != "" {
		db = db.Where("level_id = ?", params.LevelID)
	}
	if params.LocationID != "" {
		db = db.Where("location_id = ?", params.LocationID)
	}

	if params.SubcategoryID != "" {
		db = db.Where("location_id = ?", params.LocationID)
	}

	// sort
	switch params.Sort {
	case "title_asc":
		db = db.Order("title asc")
	case "title_desc":
		db = db.Order("title desc")
	case "created_asc":
		db = db.Order("created_at asc")
	case "created_desc":
		db = db.Order("created_at desc")
	default:
		db = db.Order("created_at desc")
	}

	// pagination
	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Limit(params.Limit).Offset(offset).Find(&classes).Error; err != nil {
		return nil, 0, err
	}

	return classes, count, nil
}

func (r *classRepository) FindGalleriesByClassID(classID uuid.UUID) ([]models.ClassGallery, error) {
	var galleries []models.ClassGallery
	if err := r.db.Where("class_id = ?", classID).Find(&galleries).Error; err != nil {
		return nil, err
	}
	return galleries, nil
}
