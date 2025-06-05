package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassScheduleRepository interface {
	DeleteClassSchedule(id string) error
	IncrementBooked(scheduleID uuid.UUID) error
	GetClassSchedules() ([]models.ClassSchedule, error)
	GetClassByID(id uuid.UUID) (*models.Class, error)
	HasActiveBooking(scheduleID uuid.UUID) (bool, error)
	CreateClassSchedule(schedule *models.ClassSchedule) error
	UpdateClassSchedule(schedule *models.ClassSchedule) error
	GetClassScheduleByID(id string) (*models.ClassSchedule, error)
	GetInstructorWithProfileByID(id uuid.UUID) (*models.Instructor, error)
	GetClassSchedulesWithFilter(filter dto.ClassScheduleQueryParam) ([]models.ClassSchedule, error)

	// instructor
	GetInstructorByUserID(userID uuid.UUID) (*models.Instructor, error)
	CloseScheduleWithCode(scheduleID uuid.UUID, code string) error
	OpenSchedule(scheduleID uuid.UUID, schedule *models.ClassSchedule) error
	GetAttendancesByScheduleID(scheduleID string) ([]models.Booking, error)
	GetSchedulesByInstructorID(instructorID uuid.UUID, params dto.InstructorScheduleQueryParam) ([]models.ClassSchedule, int64, error)
}

type classScheduleRepository struct {
	db *gorm.DB
}

func NewClassScheduleRepository(db *gorm.DB) ClassScheduleRepository {
	return &classScheduleRepository{db}
}

func (r *classScheduleRepository) CreateClassSchedule(schedule *models.ClassSchedule) error {
	return r.db.Create(schedule).Error
}

func (r *classScheduleRepository) UpdateClassSchedule(schedule *models.ClassSchedule) error {
	return r.db.Save(schedule).Error
}

func (r *classScheduleRepository) DeleteClassSchedule(id string) error {
	return r.db.Delete(&models.ClassSchedule{}, "id = ?", id).Error
}

func (r *classScheduleRepository) GetClassScheduleByID(id string) (*models.ClassSchedule, error) {
	var schedule models.ClassSchedule
	if err := r.db.
		First(&schedule, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *classScheduleRepository) GetClassSchedules() ([]models.ClassSchedule, error) {
	var schedules []models.ClassSchedule
	if err := r.db.
		Order("date asc").
		Order("start_hour asc").
		Order("start_minute asc").
		Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *classScheduleRepository) GetClassSchedulesWithFilter(filter dto.ClassScheduleQueryParam) ([]models.ClassSchedule, error) {
	var schedules []models.ClassSchedule
	db := r.db.
		Order("start_hour asc").
		Order("start_minute asc")

	if filter.StartDate != "" {
		if start, err := time.Parse("2006-01-02", filter.StartDate); err == nil {
			db = db.Where("date >= ?", start)
		}
	}
	if filter.EndDate != "" {
		if end, err := time.Parse("2006-01-02", filter.EndDate); err == nil {
			db = db.Where("date <= ?", end)
		}
	}

	if err := db.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *classScheduleRepository) IncrementBooked(scheduleID uuid.UUID) error {
	return r.db.Model(&models.ClassSchedule{}).
		Where("id = ?", scheduleID).
		Update("booked", gorm.Expr("booked + 1")).Error
}
func (r *classScheduleRepository) HasActiveBooking(scheduleID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.Booking{}).
		Where("class_schedule_id = ? AND status = ?", scheduleID, "booked").
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *classScheduleRepository) GetClassByID(id uuid.UUID) (*models.Class, error) {
	var class models.Class
	if err := r.db.Unscoped().Preload("Location").First(&class, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *classScheduleRepository) GetInstructorWithProfileByID(id uuid.UUID) (*models.Instructor, error) {
	var instructor models.Instructor
	if err := r.db.
		Preload("User.Profile", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		First(&instructor, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &instructor, nil
}

func (r *classScheduleRepository) GetInstructorByUserID(userID uuid.UUID) (*models.Instructor, error) {
	var instructor models.Instructor
	if err := r.db.Where("user_id = ?", userID).Find(&instructor).Error; err != nil {
		return nil, err
	}
	return &instructor, nil
}

func (r *classScheduleRepository) GetSchedulesByInstructorID(instructorID uuid.UUID, params dto.InstructorScheduleQueryParam) ([]models.ClassSchedule, int64, error) {
	var schedules []models.ClassSchedule
	var count int64

	db := r.db.Model(&models.ClassSchedule{}).
		Where("instructor_id = ? AND booked > 0", instructorID)

	if params.Status == "upcoming" {
		db = db.Where(`
		TIMESTAMP(class_schedules.date, MAKETIME(class_schedules.start_hour, class_schedules.start_minute, 0)) 
		+ INTERVAL class_schedules.duration MINUTE > CONVERT_TZ(UTC_TIMESTAMP(), '+00:00', '+07:00')
	`)
	} else if params.Status == "past" {
		db = db.Where(`
		TIMESTAMP(class_schedules.date, MAKETIME(class_schedules.start_hour, class_schedules.start_minute, 0)) 
		+ INTERVAL class_schedules.duration MINUTE <= CONVERT_TZ(UTC_TIMESTAMP(), '+00:00', '+07:00')
	`)
	}

	// Sorting
	switch params.Sort {
	case "name_asc":
		db = db.Order("class_name ASC")
	case "name_desc":
		db = db.Order("class_name DESC")
	case "date_asc":
		db = db.Order("date ASC").Order("start_hour ASC").Order("start_minute ASC")
	case "date_desc":
		db = db.Order("date DESC").Order("start_hour ASC").Order("start_minute ASC")
	default:
		db = db.Order("date DESC").Order("start_hour ASC").Order("start_minute ASC")
	}
	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// Ambil data
	if err := db.Limit(params.Limit).Offset(offset).Find(&schedules).Error; err != nil {
		return nil, 0, err
	}

	return schedules, count, nil
}

func (r *classScheduleRepository) OpenSchedule(scheduleID uuid.UUID, schedule *models.ClassSchedule) error {
	return r.db.Model(&models.ClassSchedule{}).
		Where("id = ?", scheduleID).
		Updates(map[string]any{
			"zoom_link":         schedule.ZoomLink,
			"verification_code": schedule.VerificationCode,
			"is_opened":         true,
		}).Error
}

func (r *classScheduleRepository) CloseScheduleWithCode(scheduleID uuid.UUID, code string) error {
	return r.db.Model(&models.ClassSchedule{}).
		Where("id = ?", scheduleID).
		Update("verification_code", code).Error
}

func (r *classScheduleRepository) GetAttendancesByScheduleID(scheduleID string) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.
		Preload("User.Profile").
		Preload("Attendance").
		Where("class_schedule_id = ?", scheduleID).
		Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
