package repositories

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPackageRepository interface {
	CreateUserPackage(userPackage *models.UserPackage) error
	UpdateUserPackage(userPackage *models.UserPackage) error
	GetUserPackagesByUserID(userID string) ([]models.UserPackage, error)
	GetUserPackagesByClassID(userID, classID string) ([]models.UserPackage, error)
	GetUserPackagesByPackageIDs(packageIDs []uuid.UUID) ([]models.UserPackage, error)
	GetActiveUserPackages(userID, packageID string, result *models.UserPackage) error
	GetUserPackages(userID string, params dto.PackageQueryParam) ([]models.UserPackage, int64, error)
}

type userPackageRepository struct {
	db *gorm.DB
}

func NewUserPackageRepository(db *gorm.DB) UserPackageRepository {
	return &userPackageRepository{db}
}

func (r *userPackageRepository) CreateUserPackage(userPackage *models.UserPackage) error {
	return r.db.Create(userPackage).Error
}

func (r *userPackageRepository) UpdateUserPackage(userPackage *models.UserPackage) error {
	return r.db.Save(userPackage).Error
}

func (r *userPackageRepository) GetUserPackagesByUserID(userID string) ([]models.UserPackage, error) {
	var userPackages []models.UserPackage
	err := r.db.
		Where("user_id = ?", userID).
		Order("purchased_at desc").
		Preload("Package").
		Preload("User").
		Find(&userPackages).Error
	return userPackages, err
}

func (r *userPackageRepository) GetUserPackagesByPackageIDs(packageIDs []uuid.UUID) ([]models.UserPackage, error) {
	var userPackages []models.UserPackage
	err := r.db.
		Preload("Package").
		Where("package_id IN ?", packageIDs).
		Find(&userPackages).Error
	return userPackages, err
}

func (r *userPackageRepository) GetActiveUserPackages(userID, packageID string, result *models.UserPackage) error {
	err := r.db.
		Where("user_id = ? AND package_id = ? AND expired_at > ?", userID, packageID, time.Now()).
		Order("purchased_at desc").
		First(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}

func (r *userPackageRepository) GetUserPackages(userID string, params dto.PackageQueryParam) ([]models.UserPackage, int64, error) {
	var userPackages []models.UserPackage
	var count int64

	db := r.db.Model(&models.UserPackage{}).
		Where("user_id = ?", userID).
		Preload("Package")

	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("package_name LIKE ?", like)
	}

	switch params.Sort {
	case "created_at_asc":
		db = db.Order("created_at asc")
	case "created_at_desc":
		db = db.Order("created_at desc")
	case "expired_at_asc":
		db = db.Order("expired_at asc")
	case "expired_at_desc":
		db = db.Order("expired_at desc")
	default:
		db = db.Order("created_at desc")
	}

	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Limit(params.Limit).Offset(offset).Find(&userPackages).Error; err != nil {
		return nil, 0, err
	}

	return userPackages, count, nil
}

func (r *userPackageRepository) GetUserPackagesByClassID(userID, classID string) ([]models.UserPackage, error) {
	var userPackages []models.UserPackage

	err := r.db.
		Joins("JOIN package_classes pc ON pc.package_id = user_packages.package_id").
		Where("user_packages.user_id = ? AND pc.class_id = ?", userID, classID).
		Preload("Package.Classes").
		Find(&userPackages).Error

	return userPackages, err
}
