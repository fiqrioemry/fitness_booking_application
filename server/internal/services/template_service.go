package services

import (
	"encoding/json"
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type ScheduleTemplateService interface {
	AutoGenerateSchedules() error
	RunTemplate(id string) error
	StopTemplate(id string) error
	DeleteTemplate(id string) error
	GenerateScheduleByTemplateID(templateID string) error
	GetAllTemplates() ([]dto.ScheduleTemplateResponse, error)
	CreateScheduleTemplate(req dto.CreateScheduleTemplateRequest) (string, error)
	UpdateScheduleTemplate(id string, req dto.UpdateScheduleTemplateRequest) error
}

type scheduleTemplateService struct {
	template   repositories.ScheduleTemplateRepository
	class      repositories.ClassRepository
	instructor repositories.InstructorRepository
	schedule   repositories.ClassScheduleRepository
}

func NewScheduleTemplateService(
	template repositories.ScheduleTemplateRepository,
	class repositories.ClassRepository,
	instructor repositories.InstructorRepository,
	schedule repositories.ClassScheduleRepository,
) ScheduleTemplateService {
	return &scheduleTemplateService{template, class, instructor, schedule}
}

func (s *scheduleTemplateService) GetAllTemplates() ([]dto.ScheduleTemplateResponse, error) {
	templates, err := s.template.GetAllTemplates()
	if err != nil {
		return nil, err
	}

	var result []dto.ScheduleTemplateResponse
	for _, t := range templates {
		var days []int
		_ = json.Unmarshal(t.DayOfWeeks, &days)
		resp := dto.ScheduleTemplateResponse{
			ID:             t.ID.String(),
			ClassID:        t.ClassID.String(),
			ClassName:      t.Class.Title,
			InstructorID:   t.InstructorID.String(),
			InstructorName: t.Instructor.User.Fullname,
			DayOfWeeks:     days,
			StartHour:      t.StartHour,
			StartMinute:    t.StartMinute,
			Capacity:       t.Capacity,
			IsActive:       t.IsActive,
			EndDate:        t.EndDate.Format("2006-01-02"),
			CreatedAt:      t.CreatedAt.Format("2006-01-02"),
		}
		result = append(result, resp)
	}
	return result, nil
}

func containsInt(list []int, target int) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func (s *scheduleTemplateService) CreateScheduleTemplate(req dto.CreateScheduleTemplateRequest) (string, error) {
	now := time.Now().UTC()

	endDate, err := utils.ParseDate(req.EndDate)
	if err != nil {
		return "", err
	}

	if endDate.Before(now) {
		return "", fmt.Errorf("end date must be in the future")
	}

	class, err := s.class.GetClassByID(req.ClassID)
	if err != nil {
		return "", fmt.Errorf("class not found: %w", err)
	}

	instructor, err := s.instructor.GetInstructorByID(req.InstructorID)
	if err != nil {
		return "", fmt.Errorf("instructor not found: %w", err)
	}

	for date := now; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		weekday := int(date.Weekday())
		if !containsInt(req.DayOfWeeks, weekday) {
			continue
		}

		if err := s.CheckTemplateConflict(req.InstructorID, date, req.StartHour, req.StartMinute); err != nil {
			return "", err
		}
	}

	template := models.ScheduleTemplate{
		ID:             uuid.New(),
		ClassID:        class.ID,
		ClassName:      class.Title,
		ClassImage:     class.Image,
		InstructorID:   instructor.ID,
		Location:       class.Location.Name,
		InstructorName: instructor.User.Fullname,
		DayOfWeeks:     utils.IntSliceToJSON(req.DayOfWeeks),
		StartHour:      req.StartHour,
		StartMinute:    req.StartMinute,
		Capacity:       req.Capacity,
		IsActive:       false,
		Color:          req.Color,
		EndDate:        endDate,
	}

	err = s.template.CreateTemplate(&template)
	if err != nil {
		return "", fmt.Errorf("failed to create template: %w", err)
	}

	return template.ID.String(), nil
}

func (s *scheduleTemplateService) UpdateScheduleTemplate(id string, req dto.UpdateScheduleTemplateRequest) error {
	template, err := s.template.GetTemplateByID(id)
	if err != nil {
		return err
	}

	if template.IsActive {
		return fmt.Errorf("cannot update an active template, please stop it first")
	}

	class, err := s.class.GetClassByID(req.ClassID)
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}

	instructor, err := s.instructor.GetInstructorByID(req.InstructorID)
	if err != nil {
		return fmt.Errorf("instructor not found: %w", err)
	}

	now := time.Now().UTC()
	endDate, err := utils.ParseDate(req.EndDate)
	if err != nil {
		return err
	}

	for date := now; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		weekday := int(date.Weekday())
		if !containsInt(req.DayOfWeeks, weekday) {
			continue
		}

		if err := s.CheckTemplateConflict(req.InstructorID, date, req.StartHour, req.StartMinute); err != nil {
			return err
		}
	}

	template.ClassID = class.ID
	template.ClassName = class.Title
	template.InstructorID = instructor.ID
	template.InstructorName = instructor.User.Fullname
	template.DayOfWeeks = utils.IntSliceToJSON(req.DayOfWeeks)
	template.StartHour = req.StartHour
	template.StartMinute = req.StartMinute
	template.Capacity = req.Capacity
	template.EndDate = endDate

	if err := s.template.UpdateTemplate(template); err != nil {
		return err
	}

	return nil
}

func (s *scheduleTemplateService) DeleteTemplate(id string) error {
	return s.template.DeleteTemplate(id)
}

func (s *scheduleTemplateService) RunTemplate(id string) error {
	template, err := s.template.GetTemplateByID(id)
	if err != nil {
		return err
	}
	if template.IsActive {
		return fmt.Errorf("template is already active")
	}
	template.IsActive = true
	return s.template.UpdateTemplate(template)
}

func (s *scheduleTemplateService) StopTemplate(id string) error {
	template, err := s.template.GetTemplateByID(id)
	if err != nil {
		return err
	}
	if !template.IsActive {
		return fmt.Errorf("template is already inactive")
	}
	template.IsActive = false
	return s.template.UpdateTemplate(template)
}

func (s *scheduleTemplateService) CheckTemplateConflict(instructorID string, date time.Time, hour, minute int) error {
	id := uuid.MustParse(instructorID)
	templates, err := s.template.GetAllTemplates()
	if err != nil {
		return err
	}

	weekday := int(date.Weekday())
	newStart := time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, time.UTC)
	newEnd := newStart.Add(time.Hour)

	for _, t := range templates {
		if t.InstructorID != id {
			continue
		}
		var tplDays []int
		if err := json.Unmarshal(t.DayOfWeeks, &tplDays); err != nil {
			continue
		}
		if !utils.IsDayMatched(weekday, tplDays) {
			continue
		}

		tplStart := time.Date(date.Year(), date.Month(), date.Day(), t.StartHour, t.StartMinute, 0, 0, time.UTC)
		tplEnd := tplStart.Add(time.Hour)

		if newStart.Before(tplEnd) && tplStart.Before(newEnd) {
			return fmt.Errorf("instructor %s is already booked on %s at %02d:%02d (from template)",
				t.InstructorName, date.Format("2006-01-02"), hour, minute)
		}
	}

	return nil
}

func (s *scheduleTemplateService) GenerateScheduleByTemplateID(templateID string) error {
	template, err := s.template.GetTemplateByID(templateID)
	if err != nil {
		return fmt.Errorf("failed to fetch template: %w", err)
	}

	if !template.IsActive {
		return fmt.Errorf("template is not active")
	}

	var days []int
	if err := json.Unmarshal(template.DayOfWeeks, &days); err != nil {
		return fmt.Errorf("failed to parse days of week: %w", err)
	}

	today := time.Now().UTC().Truncate(24 * time.Hour)
	end := today.AddDate(0, 1, 0)

	var hasSuccess bool
	var errors []string

	for date := today; !date.After(end); date = date.AddDate(0, 0, 1) {
		if !utils.IsDayMatched(int(date.Weekday()), days) {
			continue
		}

		instructorID := template.InstructorID.String()
		if err := s.CheckTemplateConflict(instructorID, date, template.StartHour, template.StartMinute); err != nil {
			errors = append(errors, fmt.Sprintf("conflict on %s: %v", date.Format("2006-01-02"), err))
			continue
		}
		schedule := models.ClassSchedule{
			ID:             uuid.New(),
			ClassID:        template.ClassID,
			ClassName:      template.ClassName,
			ClassImage:     template.ClassImage,
			InstructorID:   template.InstructorID,
			InstructorName: template.InstructorName,
			Location:       template.Location,
			Capacity:       template.Capacity,
			Color:          template.Color,
			Date:           date,
			StartHour:      template.StartHour,
			StartMinute:    template.StartMinute,
		}

		if err := s.schedule.CreateClassSchedule(&schedule); err != nil {
			errors = append(errors, fmt.Sprintf("failed on %s: %v", date.Format("2006-01-02"), err))
			continue
		}

		hasSuccess = true
	}

	if !hasSuccess {
		return fmt.Errorf("failed to generate any schedule: %v", errors)
	}

	now := time.Now().UTC()
	template.LastGeneratedAt = &now
	if err := s.template.UpdateTemplate(template); err != nil {
		return fmt.Errorf("schedule generated, but failed to update LastGeneratedAt: %w", err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("partial success: %v", errors)
	}

	return nil
}

// for cron job
func (s *scheduleTemplateService) AutoGenerateSchedules() error {
	templates, err := s.template.GetActiveTemplates()
	if err != nil {
		return fmt.Errorf("failed to fetch templates: %w", err)
	}

	if len(templates) == 0 {
		return fmt.Errorf("no active schedule templates found")
	}

	var anySuccess bool
	var errs []string

	for _, template := range templates {
		err := s.GenerateScheduleByTemplateID(template.ID.String())
		if err != nil {
			errs = append(errs, fmt.Sprintf("template %s: %v", template.ID.String(), err))
		} else {
			anySuccess = true
		}
	}

	if !anySuccess {
		return fmt.Errorf("no schedules generated: %v", errs)
	}

	if len(errs) > 0 {
		return fmt.Errorf("some schedules generated, with errors: %v", errs)
	}

	return nil
}
