package services

import (
	"encoding/json"
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
	"server/pkg/utils"
	"time"

	"slices"

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

	// conflict check
	CheckInstructorConflict(ID string, date time.Time, hour, minute int) error
	CheckScheduleConflict(ID string, date time.Time, hour, minute int, schedules []models.ClassSchedule) error
	CheckTemplateConflict(ID string, date time.Time, hour, minute int, templates []models.ScheduleTemplate) error
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
		return nil, customErr.NewNotFound("no templates found")
	}

	var result []dto.ScheduleTemplateResponse
	for _, t := range templates {
		var days []int
		_ = json.Unmarshal(t.DayOfWeeks, &days)
		resp := dto.ScheduleTemplateResponse{
			ID:             t.ID.String(),
			ClassID:        t.ClassID.String(),
			ClassName:      t.ClassName,
			InstructorID:   t.InstructorID.String(),
			InstructorName: t.InstructorName,
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
	return slices.Contains(list, target)
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
		return "", customErr.NewNotFound("class not found")
	}

	instructor, err := s.instructor.GetInstructorByID(req.InstructorID)
	if err != nil {
		return "", customErr.NewNotFound(fmt.Sprintf("instructor not found: %v", err))
	}

	templates, err := s.template.GetAllTemplates()
	if err != nil {
		return "", customErr.NewInternal("failed to fetch schedule templates", err)
	}

	for date := now; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		weekday := int(date.Weekday())
		if !containsInt(req.DayOfWeeks, weekday) {
			continue
		}

		if err := s.CheckTemplateConflict(req.InstructorID, date, req.StartHour, req.StartMinute, templates); err != nil {
			return "", customErr.NewConflict(err.Error())
		}
	}

	schedules, err := s.schedule.GetClassSchedules()
	if err != nil {
		return "", customErr.NewInternal("failed to fetch class schedules", err)
	}

	for date := now; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		weekday := int(date.Weekday())
		if !containsInt(req.DayOfWeeks, weekday) {
			continue
		}

		if err := s.CheckScheduleConflict(req.InstructorID, date, req.StartHour, req.StartMinute, schedules); err != nil {
			return "", customErr.NewConflict(err.Error())
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
		return "", customErr.NewInternal("failed to create template", err)
	}

	return template.ID.String(), nil
}

func (s *scheduleTemplateService) UpdateScheduleTemplate(id string, req dto.UpdateScheduleTemplateRequest) error {
	template, err := s.template.GetTemplateByID(id)
	if err != nil {
		return customErr.NewNotFound(fmt.Sprintf("template not found: %v", err))
	}
	if template.IsActive {
		return customErr.NewBadRequest("cannot update an active template, please stop it first")
	}

	endDate, err := utils.ParseDate(req.EndDate)
	if err != nil {
		return customErr.NewBadRequest("invalid end date format")
	}

	needsConflictCheck := false

	// only if class changed
	if req.ClassID != template.ClassID.String() {
		class, err := s.class.GetClassByID(req.ClassID)
		if err != nil {
			return customErr.NewNotFound("class not found")
		}
		template.ClassID = class.ID
		template.ClassName = class.Title
		template.ClassImage = class.Image
		needsConflictCheck = true
	}

	// only if instructor changed
	if req.InstructorID != template.InstructorID.String() {
		instructor, err := s.instructor.GetInstructorByID(req.InstructorID)
		if err != nil {
			return customErr.NewNotFound("instructor not found")
		}
		template.InstructorID = instructor.ID
		template.InstructorName = instructor.User.Fullname
		needsConflictCheck = true
	}

	// only if schedule rules changed
	if req.StartHour != template.StartHour || req.StartMinute != template.StartMinute ||
		!utils.IntSliceEqual(req.DayOfWeeks, utils.JSONToIntSlice(template.DayOfWeeks)) {
		needsConflictCheck = true
	}

	if needsConflictCheck {
		templates, err := s.template.GetAllTemplates()
		if err != nil {
			return customErr.NewInternal("failed to fetch schedule templates", err)
		}
		schedules, err := s.schedule.GetClassSchedules()
		if err != nil {
			return customErr.NewInternal("failed to fetch class schedules", err)
		}

		now := time.Now().UTC()
		for date := now; !date.After(endDate); date = date.AddDate(0, 0, 1) {
			weekday := int(date.Weekday())
			if !containsInt(req.DayOfWeeks, weekday) {
				continue
			}
			if err := s.CheckTemplateConflict(req.InstructorID, date, req.StartHour, req.StartMinute, templates); err != nil {
				return customErr.NewConflict(err.Error())
			}
			if err := s.CheckScheduleConflict(req.InstructorID, date, req.StartHour, req.StartMinute, schedules); err != nil {
				return customErr.NewConflict(err.Error())
			}
		}
	}

	// always update these regardless of conflict
	template.EndDate = endDate
	template.Capacity = req.Capacity
	template.StartHour = req.StartHour
	template.StartMinute = req.StartMinute
	template.DayOfWeeks = utils.IntSliceToJSON(req.DayOfWeeks)

	if err := s.template.UpdateTemplate(template); err != nil {
		return customErr.NewInternal("failed to update template", err)
	}

	return nil
}

func (s *scheduleTemplateService) DeleteTemplate(id string) error {
	err := s.template.DeleteTemplate(id)
	if err != nil {
		return customErr.NewInternal("failed to delete template", err)
	}
	return nil
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
		return customErr.NewNotFound("template not found")
	}
	if !template.IsActive {
		return customErr.NewBadRequest("template is already inactive")
	}

	template.IsActive = false

	err = s.template.UpdateTemplate(template)
	if err != nil {
		return customErr.NewInternal("failed to update template", err)
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

	var hasSuccess bool
	var errors []string

	loc, _ := time.LoadLocation("Asia/Jakarta")
	today := time.Now().In(loc).Truncate(24 * time.Hour)
	end := today.AddDate(0, 1, 0)

	// Start from tomorrow
	for date := today.AddDate(0, 0, 1); !date.After(end); date = date.AddDate(0, 0, 1) {
		fmt.Printf("🔍 Checking date: %s (Weekday: %d)\n", date.Format("2006-01-02"), date.Weekday())

		if !utils.IsDayMatched(int(date.Weekday()), days) {
			fmt.Println("⛔️ Skipped: not in template day list")
			continue
		}

		// log additional info
		fmt.Println("✅ Included: generating schedule")

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
			errMsg := fmt.Sprintf("❌ Failed on %s: %v", date.Format("2006-01-02"), err)
			fmt.Println(errMsg)
			errors = append(errors, errMsg)
			continue
		}

		fmt.Printf("📌 Created schedule for %s\n", date.Format("2006-01-02"))
		hasSuccess = true
	}

	if !hasSuccess {
		return fmt.Errorf("failed to generate any schedule: %v", errors)
	}

	now := time.Now().In(loc)
	template.LastGeneratedAt = &now
	if err := s.template.UpdateTemplate(template); err != nil {
		return fmt.Errorf("schedule generated, but failed to update LastGeneratedAt: %w", err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("partial success: %v", errors)
	}

	fmt.Println("🎉 Finished generating schedules successfully.")
	return nil
}

func (s *scheduleTemplateService) CheckInstructorConflict(ID string, date time.Time, hour, minute int) error {
	schedules, err := s.schedule.GetClassSchedules()
	if err != nil {
		return customErr.NewInternal("failed to fetch class schedules", err)
	}

	err = s.CheckScheduleConflict(ID, date, hour, minute, schedules)
	if err != nil {
		return customErr.NewConflict(err.Error())
	}

	templates, err := s.template.GetAllTemplates()
	if err != nil {
		return customErr.NewInternal("failed to fetch schedule templates", err)
	}

	err = s.CheckTemplateConflict(ID, date, hour, minute, templates)
	if err != nil {
		return customErr.NewConflict(err.Error())
	}

	return nil
}

func (s *scheduleTemplateService) CheckTemplateConflict(ID string, date time.Time, hour, minute int, templates []models.ScheduleTemplate) error {
	instructorID := uuid.MustParse(ID)
	weekday := int(date.Weekday())
	newStart := time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, time.UTC)
	newEnd := newStart.Add(time.Hour)

	for _, t := range templates {
		if t.InstructorID != instructorID {
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

func (s *scheduleTemplateService) CheckScheduleConflict(ID string, date time.Time, hour, minute int, schedules []models.ClassSchedule) error {
	instructorID := uuid.MustParse(ID)

	newStart := time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, time.UTC)
	newEnd := newStart.Add(time.Hour)

	for _, s := range schedules {
		if s.InstructorID != instructorID {
			continue
		}

		if s.Date.Year() != date.Year() || s.Date.Month() != date.Month() || s.Date.Day() != date.Day() {
			continue
		}

		existStart := time.Date(s.Date.Year(), s.Date.Month(), s.Date.Day(), s.StartHour, s.StartMinute, 0, 0, time.UTC)
		existEnd := existStart.Add(time.Hour)

		if newStart.Before(existEnd) && existStart.Before(newEnd) {
			return fmt.Errorf("instructor %s is already booked on %s at %02d:%02d",
				s.InstructorName, date.Format("2006-01-02"), hour, minute)
		}
	}

	return nil
}

// for cron job
func (s *scheduleTemplateService) AutoGenerateSchedules() error {
	templates, err := s.template.GetActiveTemplates()
	if err != nil {
		return customErr.NewInternal("failed to fetch templates", err)
	}

	if len(templates) == 0 {
		return nil
	}

	for _, template := range templates {
		if err := s.GenerateScheduleByTemplateID(template.ID.String()); err != nil {
			fmt.Printf("failed to generate for template %s: %v\n", template.ID, err)
			continue
		}

	}

	return nil
}
