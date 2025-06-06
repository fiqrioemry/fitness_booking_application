package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type ScheduleTemplateRepository interface {
	DeleteTemplate(id string) error
	GetAllTemplates() ([]models.ScheduleTemplate, error)
	CreateTemplate(template *models.ScheduleTemplate) error
	GetActiveTemplates() ([]models.ScheduleTemplate, error)
	UpdateTemplate(template *models.ScheduleTemplate) error
	GetTemplateByID(id string) (*models.ScheduleTemplate, error)
}

type scheduleTemplateRepository struct {
	db *gorm.DB
}

func NewScheduleTemplateRepository(db *gorm.DB) ScheduleTemplateRepository {
	return &scheduleTemplateRepository{db}
}

func (r *scheduleTemplateRepository) DeleteTemplate(id string) error {
	return r.db.Delete(&models.ScheduleTemplate{}, "id = ?", id).Error
}

func (r *scheduleTemplateRepository) GetAllTemplates() ([]models.ScheduleTemplate, error) {
	var templates []models.ScheduleTemplate
	err := r.db.Find(&templates).Error
	return templates, err
}
func (r *scheduleTemplateRepository) GetActiveTemplates() ([]models.ScheduleTemplate, error) {
	var templates []models.ScheduleTemplate
	err := r.db.Where("is_active = ?", true).Find(&templates).Error
	return templates, err
}

func (r *scheduleTemplateRepository) UpdateTemplate(template *models.ScheduleTemplate) error {
	return r.db.Save(template).Error
}

func (r *scheduleTemplateRepository) CreateTemplate(template *models.ScheduleTemplate) error {
	return r.db.Create(template).Error
}

func (r *scheduleTemplateRepository) GetTemplateByID(id string) (*models.ScheduleTemplate, error) {
	var template models.ScheduleTemplate
	err := r.db.First(&template, "id = ?", id).Error
	return &template, err
}
