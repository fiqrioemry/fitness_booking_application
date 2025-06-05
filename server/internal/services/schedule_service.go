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

type ClassScheduleService interface {
	// admin
	DeleteClassSchedule(id string) error
	CreateClassSchedule(req dto.CreateScheduleRequest) error
	CreateRecurringSchedule(req dto.CreateRecurringScheduleRequest) error
	UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error

	// customer
	GetSchedulesWithBookingStatus(userID string) ([]dto.ClassScheduleResponse, error)
	GetClassScheduleByID(scheduleID, userID string) (*dto.ClassScheduleDetailResponse, error)
	GetSchedulesByFilter(filter dto.ClassScheduleQueryParam) ([]dto.ClassScheduleResponse, error)

	// instructor only
	OpenClassSchedule(id string, req dto.OpenClassScheduleRequest) error
	GetAttendancesForSchedule(scheduleID string) ([]dto.AttendanceWithUserResponse, error)
	GetSchedulesByInstructor(userID string, params dto.InstructorScheduleQueryParam) ([]dto.InstructorScheduleResponse, *dto.PaginationResponse, error)
}

type classScheduleService struct {
	repo            repositories.ClassScheduleRepository
	classRepo       repositories.ClassRepository
	packageRepo     repositories.PackageRepository
	bookingRepo     repositories.BookingRepository
	templateRepo    repositories.ScheduleTemplateRepository
	templateService ScheduleTemplateService
}

func NewClassScheduleService(
	repo repositories.ClassScheduleRepository,
	classRepo repositories.ClassRepository,
	packageRepo repositories.PackageRepository,
	bookingRepo repositories.BookingRepository,
	templateRepo repositories.ScheduleTemplateRepository,
	templateService ScheduleTemplateService,
) ClassScheduleService {
	return &classScheduleService{
		repo:            repo,
		classRepo:       classRepo,
		packageRepo:     packageRepo,
		bookingRepo:     bookingRepo,
		templateRepo:    templateRepo,
		templateService: templateService,
	}
}

// Service
func (s *classScheduleService) CreateClassSchedule(req dto.CreateScheduleRequest) error {
	parsedDate, err := utils.ParseDate(req.Date)
	if err != nil {
		return err
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")

	newStartLocal := time.Date(
		parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
		req.StartHour, req.StartMinute, 0, 0, loc,
	)

	if newStartLocal.Before(time.Now().In(loc)) {
		return fmt.Errorf("cannot create schedule in the past")
	}

	newStart := time.Date(
		parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
		req.StartHour, req.StartMinute, 0, 0, time.UTC,
	)
	newEnd := newStart.Add(time.Hour)

	existingSchedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return err
	}
	for _, schedule := range existingSchedules {
		if schedule.InstructorID == uuid.MustParse(req.InstructorID) && schedule.Date.Equal(parsedDate) {
			existStart := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
				schedule.StartHour, schedule.StartMinute, 0, 0, time.UTC)
			existEnd := existStart.Add(time.Hour)

			if newStart.Before(existEnd) && existStart.Before(newEnd) {
				return fmt.Errorf("instructor already booked at this time")
			}
		}
	}

	class, err := s.repo.GetClassByID(uuid.MustParse(req.ClassID))
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}

	instructor, err := s.repo.GetInstructorWithProfileByID(uuid.MustParse(req.InstructorID))
	if err != nil {
		return fmt.Errorf("instructor not found: %w", err)
	}

	schedule := models.ClassSchedule{
		ID:             uuid.New(),
		ClassID:        class.ID,
		ClassName:      class.Title,
		ClassImage:     class.Image,
		InstructorID:   instructor.ID,
		InstructorName: instructor.User.Fullname,
		Location:       class.Location.Name,
		Duration:       class.Duration,
		Capacity:       req.Capacity,
		Color:          req.Color,
		Date:           parsedDate,
		StartHour:      req.StartHour,
		StartMinute:    req.StartMinute,
	}

	return s.repo.CreateClassSchedule(&schedule)
}

func (s *classScheduleService) CreateRecurringSchedule(req dto.CreateRecurringScheduleRequest) error {
	now := time.Now().UTC()

	endDate, err := utils.ParseDate(req.EndDate)
	if err != nil {
		return err
	}

	if endDate.Before(now) {
		return fmt.Errorf("end date must be in the future")
	}

	instructorID := uuid.MustParse(req.InstructorID)

	existingSchedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return err
	}

	existingTemplates, err := s.templateRepo.GetAllTemplates()
	if err != nil {
		return err
	}

	for date := now; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		weekday := int(date.Weekday())
		if !utils.ContainsInt(req.DayOfWeeks, weekday) {
			continue
		}

		simulatedStart := time.Date(date.Year(), date.Month(), date.Day(),
			req.StartHour, req.StartMinute, 0, 0, time.UTC)
		simulatedEnd := simulatedStart.Add(time.Hour)

		for _, schedule := range existingSchedules {
			if schedule.InstructorID != instructorID {
				continue
			}
			if schedule.Date.Year() != date.Year() ||
				schedule.Date.Month() != date.Month() ||
				schedule.Date.Day() != date.Day() {
				continue
			}

			existStart := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
				schedule.StartHour, schedule.StartMinute, 0, 0, time.UTC)
			existEnd := existStart.Add(time.Hour)

			if simulatedStart.Before(existEnd) && existStart.Before(simulatedEnd) {
				return fmt.Errorf("instructor %s is already booked on %s at %02d:%02d (existing schedule)",
					schedule.InstructorName, date.Format("2006-01-02"), req.StartHour, req.StartMinute)
			}
		}

		for _, tpl := range existingTemplates {
			if tpl.InstructorID != instructorID {
				continue
			}
			var tplDays []int
			if err := json.Unmarshal(tpl.DayOfWeeks, &tplDays); err != nil {
				continue
			}
			if !utils.ContainsInt(tplDays, weekday) {
				continue
			}

			tplStart := time.Date(date.Year(), date.Month(), date.Day(),
				tpl.StartHour, tpl.StartMinute, 0, 0, time.UTC)
			tplEnd := tplStart.Add(time.Hour)

			if simulatedStart.Before(tplEnd) && tplStart.Before(simulatedEnd) {
				return fmt.Errorf("instructor %s is already booked on %s at %02d:%02d (template schedule)",
					tpl.InstructorName, date.Format("2006-01-02"), req.StartHour, req.StartMinute)
			}
		}
	}

	class, err := s.repo.GetClassByID(uuid.MustParse(req.ClassID))
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}

	instructor, err := s.repo.GetInstructorWithProfileByID(instructorID)
	if err != nil {
		return fmt.Errorf("instructor not found: %w", err)
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

	err = s.templateRepo.CreateTemplate(&template)
	if err != nil {
		return fmt.Errorf("failed to create template: %w", err)
	}

	return s.templateService.GenerateScheduleByTemplateID(template.ID.String())
}

func (s *classScheduleService) UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error {
	const action = "update"
	const service = "class_schedule"

	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return fmt.Errorf("schedule not found")
	}

	parsedDate, err := utils.ParseDate(req.Date)
	if err != nil {
		return fmt.Errorf("invalid date format")
	}

	newStart := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), req.StartHour, req.StartMinute, 0, 0, time.UTC)
	if newStart.Before(time.Now().UTC()) {
		return fmt.Errorf("cannot update schedule to a past time")
	}

	newEnd := newStart.Add(time.Hour)

	existingSchedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return err
	}

	instructorID := uuid.MustParse(req.InstructorID)
	for _, other := range existingSchedules {
		if other.ID == schedule.ID {
			continue
		}
		if other.InstructorID == instructorID && other.Date.Equal(parsedDate) {
			existStart := time.Date(other.Date.Year(), other.Date.Month(), other.Date.Day(), other.StartHour, other.StartMinute, 0, 0, time.UTC)
			existEnd := existStart.Add(time.Hour)

			if newStart.Before(existEnd) && existStart.Before(newEnd) {
				return fmt.Errorf("instructor already booked at this time")
			}
		}
	}

	if req.Capacity < schedule.Booked {
		return fmt.Errorf("capacity cannot be less than booked participant (%d)", schedule.Booked)
	}

	classID := uuid.MustParse(req.ClassID)
	class, err := s.repo.GetClassByID(classID)
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}

	instructor, err := s.repo.GetInstructorWithProfileByID(instructorID)
	if err != nil {
		return fmt.Errorf("instructor not found: %w", err)
	}

	schedule.Color = req.Color
	schedule.Date = parsedDate
	schedule.ClassID = class.ID
	schedule.Capacity = req.Capacity
	schedule.ClassName = class.Title
	schedule.ClassImage = class.Image
	schedule.StartHour = req.StartHour
	schedule.Duration = class.Duration
	schedule.InstructorID = instructor.ID
	schedule.StartMinute = req.StartMinute
	schedule.InstructorName = instructor.User.Fullname

	if err := s.repo.UpdateClassSchedule(schedule); err != nil {
		return err
	}
	return nil
}

func (s *classScheduleService) DeleteClassSchedule(id string) error {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return fmt.Errorf("schedule not found")
	}

	startTime := time.Date(
		schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
		schedule.StartHour, schedule.StartMinute, 0, 0, time.Local,
	)

	if startTime.Before(time.Now().UTC()) {
		return fmt.Errorf("cannot delete past or ongoing class schedule")
	}

	isBooked, err := s.repo.HasActiveBooking(schedule.ID)
	if err != nil {
		return fmt.Errorf("failed to check booking: %w", err)
	}
	if isBooked {
		return fmt.Errorf("cannot delete schedule with active bookings")
	}

	return s.repo.DeleteClassSchedule(id)
}

func (s *classScheduleService) GetClassScheduleByID(scheduleID, userID string) (*dto.ClassScheduleDetailResponse, error) {
	schedule, err := s.repo.GetClassScheduleByID(scheduleID)
	if err != nil {
		return nil, err
	}

	packages, err := s.packageRepo.GetPackagesByClassID(schedule.ClassID.String())
	if err != nil {
		return nil, err
	}

	var pkgResponses []dto.PackageListResponse
	for _, p := range packages {
		pkgResponses = append(pkgResponses, dto.PackageListResponse{
			ID:    p.ID.String(),
			Name:  p.Name,
			Price: p.Price,
			Image: p.Image,
		})
	}

	isBooked, _ := s.bookingRepo.IsUserBookedSchedule(userID, scheduleID)

	return &dto.ClassScheduleDetailResponse{
		ClassScheduleResponse: dto.ClassScheduleResponse{
			ID:             schedule.ID.String(),
			ClassID:        schedule.ClassID.String(),
			ClassName:      schedule.ClassName,
			ClassImage:     schedule.ClassImage,
			InstructorID:   schedule.InstructorID.String(),
			InstructorName: schedule.InstructorName,
			Location:       schedule.Location,
			Date:           schedule.Date.Format("2006-01-02"),
			StartHour:      schedule.StartHour,
			StartMinute:    schedule.StartMinute,
			Capacity:       schedule.Capacity,
			BookedCount:    schedule.Booked,
			Color:          schedule.Color,
			Duration:       schedule.Duration,
			IsBooked:       isBooked,
		},
		Packages: pkgResponses,
	}, nil
}

func (s *classScheduleService) GetSchedulesByFilter(filter dto.ClassScheduleQueryParam) ([]dto.ClassScheduleResponse, error) {
	schedules, err := s.repo.GetClassSchedulesWithFilter(filter)
	if err != nil {
		return nil, err
	}

	var result []dto.ClassScheduleResponse
	for _, schedule := range schedules {
		result = append(result, dto.ClassScheduleResponse{
			ID:             schedule.ID.String(),
			ClassID:        schedule.ClassID.String(),
			ClassName:      schedule.ClassName,
			ClassImage:     schedule.ClassImage,
			InstructorID:   schedule.InstructorID.String(),
			InstructorName: schedule.InstructorName,
			Location:       schedule.Location,
			Date:           schedule.Date.Format("2006-01-02"),
			StartHour:      schedule.StartHour,
			StartMinute:    schedule.StartMinute,
			Capacity:       schedule.Capacity,
			Duration:       schedule.Duration,
			BookedCount:    schedule.Booked,
			Color:          schedule.Color,
			IsBooked:       false,
		})
	}

	return result, nil
}

func (s *classScheduleService) GetSchedulesWithBookingStatus(userID string) ([]dto.ClassScheduleResponse, error) {
	schedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return nil, err
	}

	var result []dto.ClassScheduleResponse
	for _, schedule := range schedules {
		isBooked, _ := s.bookingRepo.IsUserBookedSchedule(userID, schedule.ID.String())

		result = append(result, dto.ClassScheduleResponse{
			ID:             schedule.ID.String(),
			ClassID:        schedule.ClassID.String(),
			ClassName:      schedule.ClassName,
			ClassImage:     schedule.ClassImage,
			InstructorID:   schedule.InstructorID.String(),
			InstructorName: schedule.InstructorName,
			Date:           schedule.Date.Format("2006-01-02"),
			StartHour:      schedule.StartHour,
			Location:       schedule.Location,
			StartMinute:    schedule.StartMinute,
			Duration:       schedule.Duration,
			Capacity:       schedule.Capacity,
			BookedCount:    schedule.Booked,
			Color:          schedule.Color,
			IsBooked:       isBooked,
		})
	}

	return result, nil
}

func (s *classScheduleService) GetSchedulesByInstructor(userID string, params dto.InstructorScheduleQueryParam) ([]dto.InstructorScheduleResponse, *dto.PaginationResponse, error) {
	ID := uuid.MustParse(userID)

	instructor, err := s.repo.GetInstructorByUserID(ID)
	if err != nil {
		return nil, nil, fmt.Errorf("instructor not found: %w with %s", err, instructor.ID)
	}

	schedules, total, err := s.repo.GetSchedulesByInstructorID(instructor.ID, params)
	if err != nil {
		return nil, nil, err
	}

	var results []dto.InstructorScheduleResponse
	for _, schedule := range schedules {
		results = append(results, dto.InstructorScheduleResponse{
			ID:               schedule.ID.String(),
			ClassID:          schedule.ClassID.String(),
			ClassName:        schedule.ClassName,
			ClassImage:       schedule.ClassImage,
			InstructorID:     schedule.InstructorID.String(),
			InstructorName:   schedule.InstructorName,
			Location:         schedule.Location,
			StartHour:        schedule.StartHour,
			StartMinute:      schedule.StartMinute,
			Capacity:         schedule.Capacity,
			Duration:         schedule.Duration,
			BookedCount:      schedule.Booked,
			IsOpened:         schedule.IsOpened,
			Date:             schedule.Date.Format("2006-01-02"),
			ZoomLink:         utils.EmptyString(schedule.ZoomLink),
			VerificationCode: utils.EmptyString(schedule.VerificationCode),
		})
	}
	pagination := utils.Paginate(total, params.Page, params.Limit)
	return results, pagination, nil
}

func (s *classScheduleService) OpenClassSchedule(id string, req dto.OpenClassScheduleRequest) error {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return fmt.Errorf("schedule not found")
	}
	if schedule.IsOpened {
		return fmt.Errorf("schedule already opened")
	}

	if req.ZoomLink != "" {
		schedule.ZoomLink = &req.ZoomLink
	}

	schedule.VerificationCode = &req.VerificationCode

	return s.repo.OpenSchedule(schedule.ID, schedule)
}

func (s *classScheduleService) GetAttendancesForSchedule(scheduleID string) ([]dto.AttendanceWithUserResponse, error) {
	bookings, err := s.repo.GetAttendancesByScheduleID(scheduleID)
	if err != nil {
		return nil, err
	}

	var result []dto.AttendanceWithUserResponse
	for _, b := range bookings {
		attendance := b.Attendance
		user := b.User

		resp := dto.AttendanceWithUserResponse{
			Fullname:   user.Fullname,
			Avatar:     user.Avatar,
			Email:      user.Email,
			Status:     attendance.Status,
			CheckedIn:  attendance.CheckedIn,
			CheckedOut: attendance.CheckedOut,
		}

		if attendance.CheckedAt != nil && !attendance.CheckedAt.IsZero() {
			resp.CheckedAt = attendance.CheckedAt.Format(time.RFC3339)
		}
		if attendance.VerifiedAt != nil && !attendance.VerifiedAt.IsZero() {
			resp.VerifiedAt = attendance.VerifiedAt.Format(time.RFC3339)
		}

		result = append(result, resp)
	}

	return result, nil
}
