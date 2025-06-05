package services

import (
	"errors"
	"fmt"
	"log"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
	"server/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingService interface {
	MarkAbsentBookings() error
	CheckedInClassSchedule(userID, bookingID string) error
	CreateBooking(userID, packageID, scheduleID string) error
	GetBookingDetail(userID, bookingID string) (*dto.BookingDetailResponse, error)
	CheckoutClassSchedule(userID, bookingID string, req dto.ValidateCheckoutRequest) error
	GetBookingByUser(userID string, params dto.BookingQueryParam) ([]dto.BookingResponse, *dto.PaginationResponse, error)
}

type bookingService struct {
	db           *gorm.DB
	booking      repositories.BookingRepository
	pkg          repositories.PackageRepository
	notification NotificationService
	userPkg      repositories.UserPackageRepository
	schedule     repositories.ClassScheduleRepository
}

func NewBookingService(db *gorm.DB, booking repositories.BookingRepository, pkg repositories.PackageRepository, notification NotificationService, userPkg repositories.UserPackageRepository, schedule repositories.ClassScheduleRepository) BookingService {
	return &bookingService{
		db:           db,
		booking:      booking,
		pkg:          pkg,
		notification: notification,
		userPkg:      userPkg,
		schedule:     schedule,
	}
}

func (s *bookingService) CreateBooking(userID, packageID, scheduleID string) error {
	schedule, err := s.schedule.GetClassScheduleByID(scheduleID)
	if err != nil {
		return customErr.NewNotFound("Class schedule not found")
	}

	var userPackage models.UserPackage
	err = s.userPkg.GetActiveUserPackages(userID, packageID, &userPackage)
	if err != nil {
		return customErr.NewNotFound("You don’t have an active package for this class")
	}
	if userPackage.RemainingCredit <= 0 {
		return customErr.NewConflict("Not enough credit")
	}

	count, err := s.booking.CountBookingBySchedule(schedule.ID.String())
	if err != nil {
		return customErr.NewInternal("Failed to count schedule bookings", err)
	}
	if int(count) >= schedule.Capacity {
		return customErr.NewConflict("Class schedule is full")
	}

	bookingID := uuid.New()
	attendanceID := uuid.New()

	err = s.db.Transaction(func(tx *gorm.DB) error {
		booking := models.Booking{
			ID:              bookingID,
			UserID:          uuid.MustParse(userID),
			ClassScheduleID: schedule.ID,
			Status:          "booked",
		}
		if err := tx.Create(&booking).Error; err != nil {
			return err
		}

		attendance := models.Attendance{
			ID:        attendanceID,
			BookingID: booking.ID,
		}
		if err := tx.Create(&attendance).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.UserPackage{}).
			Where("id = ?", userPackage.ID).
			Update("remaining_credit", gorm.Expr("remaining_credit - ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.ClassSchedule{}).
			Where("id = ?", schedule.ID).
			Update("booked", gorm.Expr("booked + 1")).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return customErr.NewInternal("Failed to create booking", err)
	}

	// Kirim notifikasi
	payload := dto.NotificationEvent{
		UserID: bookingID.String(),
		Type:   "system_message",
		Title:  "Class Booked Successfully",
		Message: fmt.Sprintf(
			"You have successfully booked the class \"%s\" on %s at %02d:%02d. 1 credit has been deducted from your package.",
			schedule.ClassName,
			schedule.Date.Format("January 2, 2006"),
			schedule.StartHour,
			schedule.StartMinute,
		),
	}
	if err := s.notification.SendToUser(payload); err != nil {
		log.Printf("Failed sending notification to user %s: %v\n", payload.UserID, err)
	}

	return nil
}

func (s *bookingService) GetBookingByUser(userID string, params dto.BookingQueryParam) ([]dto.BookingResponse, *dto.PaginationResponse, error) {
	bookings, total, err := s.booking.GetBookingsByUserID(userID, params)
	if err != nil {
		return nil, nil, customErr.NewNotFound("booking not found")
	}

	var result []dto.BookingResponse
	for _, b := range bookings {
		schedule := b.ClassSchedule

		result = append(result, dto.BookingResponse{
			ID:             b.ID.String(),
			BookingStatus:  b.Status,
			ClassID:        schedule.ClassID.String(),
			ClassName:      schedule.ClassName,
			ClassImage:     schedule.ClassImage,
			InstructorName: schedule.InstructorName,
			StartHour:      schedule.StartHour,
			StartMinute:    schedule.StartMinute,
			Duration:       schedule.Duration,
			Location:       schedule.Location,
			IsOpened:       schedule.IsOpened,
			Date:           schedule.Date.Format("2006-01-02"),
			BookedAt:       b.CreatedAt.Format(time.RFC3339),
		})
	}
	pagination := utils.Paginate(total, params.Page, params.Limit)
	return result, pagination, nil
}

func (s *bookingService) GetBookingDetail(userID, bookingID string) (*dto.BookingDetailResponse, error) {
	booking, err := s.booking.GetBookingByID(userID, bookingID)
	if err != nil {
		return nil, customErr.NewNotFound("booking not found")
	}

	attendance := booking.Attendance
	schedule := booking.ClassSchedule

	res := &dto.BookingDetailResponse{
		ID:               booking.ID.String(),
		ScheduleID:       schedule.ID.String(),
		ClassID:          schedule.ClassID.String(),
		ClassName:        schedule.ClassName,
		ClassImage:       schedule.ClassImage,
		InstructorName:   schedule.InstructorName,
		Date:             schedule.Date.Format("2006-01-02"),
		StartHour:        schedule.StartHour,
		StartMinute:      schedule.StartMinute,
		Duration:         schedule.Duration,
		CheckedIn:        attendance.CheckedIn,
		CheckedOut:       attendance.CheckedOut,
		IsOpened:         schedule.IsOpened,
		IsReviewed:       attendance.IsReviewed,
		AttendanceStatus: attendance.Status,
		CheckedAt:        "",
		VerifiedAt:       "",
	}

	if schedule.ZoomLink != nil {
		res.ZoomLink = *schedule.ZoomLink
	}

	if attendance.CheckedIn && attendance.CheckedAt != nil {
		res.CheckedAt = attendance.CheckedAt.Format(time.RFC3339)
	}
	if attendance.CheckedOut && attendance.VerifiedAt != nil {
		res.VerifiedAt = attendance.VerifiedAt.Format(time.RFC3339)

	}

	return res, nil
}

func (s *bookingService) CheckedInClassSchedule(userID, bookingID string) error {
	booking, err := s.booking.GetBookingByID(userID, bookingID)
	if err != nil {
		return customErr.NewNotFound("booking not found")
	}

	if !booking.ClassSchedule.IsOpened {
		return customErr.NewForbidden("Class schedule is not opened yet")
	}

	attendance := booking.Attendance

	if attendance.CheckedIn {
		return customErr.NewConflict("You have already checked in to this class")
	}

	now := time.Now().UTC()
	attendance.CheckedIn = true
	attendance.Status = "entered"
	attendance.CheckedAt = &now

	err = s.booking.UpdateAttendance(&attendance)
	if err != nil {
		return customErr.NewInternal("Failed to update attendance", err)
	}

	return nil
}

func (s *bookingService) CheckoutClassSchedule(userID, bookingID string, req dto.ValidateCheckoutRequest) error {
	booking, err := s.booking.GetBookingByID(userID, bookingID)
	if err != nil {
		return errors.New("booking not found")
	}
	attendance := booking.Attendance

	if attendance.CheckedOut {
		return customErr.NewConflict("You have already checked out from this class")
	}

	if booking.ClassSchedule.VerificationCode == nil || *booking.ClassSchedule.VerificationCode != req.VerificationCode {
		return customErr.NewBadRequest("Invalid verification code")
	}

	now := time.Now().UTC()
	attendance.CheckedOut = true
	attendance.Status = "attended"
	attendance.VerifiedAt = &now

	err = s.booking.UpdateAttendance(&attendance)
	if err != nil {
		return customErr.NewInternal("Failed to update attendance", err)
	}
	return nil
}

// ** buat cron job
func (s *bookingService) MarkAbsentBookings() error {
	now := time.Now().UTC()

	bookings, err := s.booking.GetAllBookedWithScheduleAndClass()
	if err != nil {
		return err
	}

	var totalMarked int
	for _, b := range bookings {
		if b.Status != "booked" {
			continue
		}

		schedule := b.ClassSchedule

		startTime := time.Date(
			schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
			schedule.StartHour, schedule.StartMinute, 0, 0, time.UTC,
		)
		endTime := startTime.Add(time.Duration(schedule.Duration) * time.Minute)

		if now.After(endTime) {
			exists, err := s.booking.CheckAttendanceExists(b.ID)
			if err != nil {
				log.Printf("Failed to check attendance for booking %s: %v\n", b.ID, err)
				continue
			}

			if !exists {
				attendance := &models.Attendance{
					ID:        uuid.New(),
					BookingID: b.ID,
					Status:    "absent",
				}

				if err := s.booking.CreateAttendance(attendance); err != nil {
					log.Printf("Failed to create absent attendance for booking %s: %v\n", b.ID, err)
				} else {
					totalMarked++
				}
			}
		}
	}

	return nil
}
