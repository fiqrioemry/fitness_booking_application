package repositories

import (
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(review *models.Review) error
	GetReviewsByClassID(classID string) ([]models.Review, error)
	GetAverageRatingByInstructorID(instructorID uuid.UUID) (float64, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db}
}

func (r *reviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *reviewRepository) GetReviewsByClassID(classID string) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Preload("User").Preload("Class").
		Where("class_id = ?", classID).
		Order("created_at desc").
		Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetAverageRatingByInstructorID(instructorID uuid.UUID) (float64, error) {
	var avgRating float64
	err := r.db.
		Table("reviews").
		Select("AVG(rating)").
		Joins("JOIN classes ON reviews.class_id = classes.id").
		Joins("JOIN class_schedules ON classes.id = class_schedules.class_id").
		Where("class_schedules.instructor_id = ?", instructorID).
		Scan(&avgRating).Error
	return avgRating, err
}
