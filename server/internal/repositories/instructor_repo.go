package repositories

import (
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InstructorRepository interface {
	DeleteInstructor(id string) error
	GetAllInstructors() ([]models.Instructor, error)
	UpdateInstructor(instructor *models.Instructor) error
	CreateInstructor(instructor *models.Instructor) error
	GetInstructorByID(id string) (*models.Instructor, error)
	UpdateRating(instructorID uuid.UUID, rating float64) error
}

type instructorRepository struct {
	db *gorm.DB
}

func NewInstructorRepository(db *gorm.DB) InstructorRepository {
	return &instructorRepository{db}
}

func (r *instructorRepository) CreateInstructor(instructor *models.Instructor) error {
	return r.db.Create(instructor).Error
}

func (r *instructorRepository) UpdateInstructor(instructor *models.Instructor) error {
	return r.db.Save(instructor).Error
}

func (r *instructorRepository) DeleteInstructor(id string) error {
	return r.db.Delete(&models.Instructor{}, "id = ?", id).Error
}

func (r *instructorRepository) GetInstructorByID(id string) (*models.Instructor, error) {
	var instructor models.Instructor
	if err := r.db.Preload("User").First(&instructor, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &instructor, nil
}

func (r *instructorRepository) GetAllInstructors() ([]models.Instructor, error) {
	var instructors []models.Instructor
	if err := r.db.Preload("User").Find(&instructors).Error; err != nil {
		return nil, err
	}
	return instructors, nil
}

func (r *instructorRepository) UpdateRating(instructorID uuid.UUID, rating float64) error {
	return r.db.Model(&models.Instructor{}).
		Where("id = ?", instructorID).
		Update("rating", rating).Error
}
