package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	UpdateUser(user *models.User) error
	GetUserByID(userID string) (*models.User, error)
	GetUserStats() (int64, int64, int64, int64, int64, error)
	FindAllUsers(params dto.UserQueryParam) ([]models.User, int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(user).Error
}

func (r *userRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Tokens").First(&user, "id = ?", id).Error
	return &user, err
}

func (r *userRepository) FindAllUsers(params dto.UserQueryParam) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	db := r.db.Model(&models.User{})

	if params.Q != "" {
		db = db.Where("email LIKE ? OR fullname LIKE ?", "%"+params.Q+"%", "%"+params.Q+"%")
	}
	if params.Role != "" && params.Role != "all" {
		db = db.Where("role = ?", params.Role)
	}

	switch params.Sort {
	case "joined_asc":
		db = db.Order("created_at asc")
	case "joined_desc":
		db = db.Order("created_at desc")
	case "email_asc":
		db = db.Order("email asc")
	case "email_desc":
		db = db.Order("email desc")
	case "name_asc":
		db = db.Order("fullname asc")
	case "name_desc":
		db = db.Order("fullname desc")
	default:
		db = db.Order("created_at desc")
	}

	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Limit(params.Limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func (r *userRepository) GetUserStats() (int64, int64, int64, int64, int64, error) {
	var total, customers, instructors, admins, newThisMonth int64
	var err error
	db := r.db.Model(&models.User{})

	db.Count(&total)

	db.Where("role = ?", "customer").Count(&customers)

	db.Where("role = ?", "instructor").Count(&instructors)

	db.Where("role = ?", "admin").Count(&admins)

	db.Where("created_at >= ?", time.Now().AddDate(0, -1, 0)).Count(&newThisMonth)

	return total, customers, instructors, admins, newThisMonth, err
}
