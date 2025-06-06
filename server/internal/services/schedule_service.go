package services

import (
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
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
	schedule    repositories.ClassScheduleRepository
	template    ScheduleTemplateService
	class       repositories.ClassRepository
	instructor  repositories.InstructorRepository
	bookingRepo repositories.BookingRepository
	packageRepo repositories.PackageRepository
}

func NewClassScheduleService(
	schedule repositories.ClassScheduleRepository,
	template ScheduleTemplateService,
	class repositories.ClassRepository,
	instructor repositories.InstructorRepository,
	bookingRepo repositories.BookingRepository,
	packageRepo repositories.PackageRepository,
) ClassScheduleService {
	return &classScheduleService{
		schedule:    schedule,
		template:    template,
		class:       class,
		instructor:  instructor,
		bookingRepo: bookingRepo,
		packageRepo: packageRepo,
	}
}

// Service
func (s *classScheduleService) CreateClassSchedule(req dto.CreateScheduleRequest) error {

	parsedDate, err := utils.ParseDate(req.Date)
	if err != nil {
		return err
	}

	if err := utils.ValidateScheduleNotInPast(parsedDate, req.StartHour, req.StartMinute); err != nil {
		return err
	}

	class, err := s.class.GetClassByID(req.ClassID)
	if err != nil {
		return customErr.NewNotFound("class not found")
	}

	instructor, err := s.instructor.GetInstructorByID(req.InstructorID)
	if err != nil {
		return customErr.NewNotFound("instructor not found")
	}

	err = s.template.CheckInstructorConflict(req.InstructorID, parsedDate, req.StartHour, req.StartMinute)
	if err != nil {
		return customErr.NewConflict(err.Error())
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

	err = s.schedule.CreateClassSchedule(&schedule)
	if err != nil {
		return customErr.NewInternal("Failed to create schedule", err)
	}

	return nil
}

func (s *classScheduleService) CreateRecurringSchedule(req dto.CreateRecurringScheduleRequest) error {
	templateReq := dto.CreateScheduleTemplateRequest{
		ClassID:      req.ClassID,
		InstructorID: req.InstructorID,
		DayOfWeeks:   req.DayOfWeeks,
		StartHour:    req.StartHour,
		StartMinute:  req.StartMinute,
		Date:         req.Date,
		Capacity:     req.Capacity,
		Color:        req.Color,
		EndDate:      req.EndDate,
	}

	templateID, err := s.template.CreateScheduleTemplate(templateReq)
	if err != nil {
		return err
	}

	err = s.template.GenerateScheduleByTemplateID(templateID)
	if err != nil {
		return err
	}
	return nil
}

func (s *classScheduleService) UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error {

	parsedDate, err := utils.ParseDate(req.Date)
	if err != nil {
		return err
	}

	if err := utils.ValidateScheduleNotInPast(parsedDate, req.StartHour, req.StartMinute); err != nil {
		return err
	}

	schedule, err := s.schedule.GetClassScheduleByID(id)
	if err != nil {
		return customErr.NewNotFound("schedule not found")
	}

	if req.Capacity < schedule.Booked {
		return fmt.Errorf("capacity cannot be less than booked participant (%d)", schedule.Booked)
	}

	// only if class changed, do checking
	if req.ClassID != schedule.ClassID.String() {
		class, err := s.class.GetClassByID(req.ClassID)
		if err != nil {
			return customErr.NewNotFound("class not found")
		}
		schedule.ClassID = class.ID
		schedule.ClassName = class.Title
		schedule.ClassImage = class.Image
		schedule.Duration = class.Duration
	}

	// only if instructor changed, do checking
	if req.InstructorID != schedule.InstructorID.String() {
		instructor, err := s.instructor.GetInstructorByID(req.InstructorID)
		if err != nil {
			return customErr.NewNotFound("instructor not found")
		}
		schedule.InstructorID = instructor.ID
		schedule.InstructorName = instructor.User.Fullname
	}

	err = s.template.CheckInstructorConflict(req.InstructorID, parsedDate, req.StartHour, req.StartMinute)
	if err != nil {
		return customErr.NewConflict(err.Error())
	}

	schedule.Color = req.Color
	schedule.Date = parsedDate
	schedule.Capacity = req.Capacity
	schedule.StartHour = req.StartHour
	schedule.StartMinute = req.StartMinute

	if err := s.schedule.UpdateClassSchedule(schedule); err != nil {
		return err
	}

	return nil
}

func (s *classScheduleService) DeleteClassSchedule(id string) error {
	schedule, err := s.schedule.GetClassScheduleByID(id)
	if err != nil {
		return customErr.NewNotFound("schedule not found")
	}

	if err := utils.ValidateScheduleNotInPast(schedule.Date, schedule.StartHour, schedule.StartMinute); err != nil {
		return err
	}

	isBooked, err := s.schedule.HasActiveBooking(schedule.ID)
	if err != nil {
		return customErr.NewInternal("failed to check active booking", err)
	}
	if isBooked {
		return customErr.NewBadRequest("cannot delete schedule with active bookings")
	}

	err = s.schedule.DeleteClassSchedule(id)
	if err != nil {
		return customErr.NewInternal("failed to delete schedule", err)
	}

	return nil
}

func (s *classScheduleService) GetClassScheduleByID(scheduleID, userID string) (*dto.ClassScheduleDetailResponse, error) {
	schedule, err := s.schedule.GetClassScheduleByID(scheduleID)
	if err != nil {
		return nil, customErr.NewNotFound("schedule not found")
	}

	packages, err := s.packageRepo.GetPackagesByClassID(schedule.ClassID.String())
	if err != nil {
		return nil, customErr.NewNotFound("no packages found")
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
	schedules, err := s.schedule.GetClassSchedulesWithFilter(filter)
	if err != nil {
		return nil, customErr.NewNotFound("no schedules found")
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
	schedules, err := s.schedule.GetClassSchedules()
	if err != nil {
		return nil, customErr.NewNotFound("no schedules found")
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

	instructor, err := s.instructor.GetInstructorByUserID(userID)
	if err != nil {
		return nil, nil, customErr.NewNotFound(fmt.Sprintf("instructor not found for user %s: %v", userID, err))
	}

	schedules, total, err := s.schedule.GetSchedulesByInstructorID(instructor.ID, params)
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
	schedule, err := s.schedule.GetClassScheduleByID(id)
	if err != nil {
		return customErr.NewNotFound("no schedule found")
	}
	if schedule.IsOpened {
		return fmt.Errorf("schedule already opened")
	}

	if req.ZoomLink != "" {
		schedule.ZoomLink = &req.ZoomLink
	}

	schedule.VerificationCode = &req.VerificationCode

	return s.schedule.OpenSchedule(schedule.ID, schedule)
}

func (s *classScheduleService) GetAttendancesForSchedule(scheduleID string) ([]dto.AttendanceWithUserResponse, error) {
	bookings, err := s.schedule.GetAttendancesByScheduleID(scheduleID)
	if err != nil {
		return nil, customErr.NewNotFound("no attendance found")
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
