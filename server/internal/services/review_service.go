package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewService interface {
	GetReviewsByClassID(classID string) ([]dto.ReviewResponse, error)
	CreateReview(userID string, bookingID string, req dto.CreateReviewRequest) error
}

type reviewService struct {
	review     repositories.ReviewRepository
	booking    repositories.BookingRepository
	instructor repositories.InstructorRepository
}

func NewReviewService(review repositories.ReviewRepository, booking repositories.BookingRepository, instructor repositories.InstructorRepository) ReviewService {
	return &reviewService{review, booking, instructor}
}

func (s *reviewService) GetReviewsByClassID(classID string) ([]dto.ReviewResponse, error) {
	reviews, err := s.review.GetReviewsByClassID(classID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewNotFound("Reviews for this class not found")
		}
		return nil, customErr.NewInternal("Failed to retrieve reviews", err)
	}

	var result []dto.ReviewResponse
	for _, r := range reviews {
		result = append(result, dto.ReviewResponse{
			ID:        r.ID.String(),
			UserName:  r.User.Fullname,
			ClassName: r.Class.Title,
			Rating:    r.Rating,
			Comment:   r.Comment,
			CreatedAt: r.CreatedAt.Format(time.RFC3339),
		})
	}
	return result, nil
}

func (s *reviewService) CreateReview(userID string, bookingID string, req dto.CreateReviewRequest) error {

	booking, err := s.booking.GetBookingByID(userID, bookingID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewNotFound("Booking not found")
		}
		return customErr.NewInternal("Failed to get booking details", err)
	}

	attendance := booking.Attendance
	if attendance.IsReviewed {
		return customErr.NewConflict("You have already submitted a review")
	}

	schedule := booking.ClassSchedule
	review := models.Review{
		UserID:  uuid.MustParse(userID),
		ClassID: schedule.ClassID,
		Rating:  req.Rating,
		Comment: req.Comment,
	}

	if err := s.review.CreateReview(&review); err != nil {
		return customErr.NewInternal("Failed to create review", err)
	}

	attendance.IsReviewed = true
	if err := s.booking.UpdateAttendance(&attendance); err != nil {
		return customErr.NewInternal("Failed to update attendance", err)
	}

	avgRating, err := s.review.GetAverageRatingByInstructorID(schedule.InstructorID)
	if err != nil {
		return customErr.NewInternal("Failed to calculate instructor average rating", err)
	}

	if err := s.instructor.UpdateRating(schedule.InstructorID, avgRating); err != nil {
		return customErr.NewInternal("Failed to update instructor rating", err)
	}

	return nil
}
