package repositories

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"

	"gorm.io/gorm"
)

type PackageRepository interface {
	DeletePackage(id string) error
	CreatePackage(pkg *models.Package) error
	UpdatePackage(pkg *models.Package) error
	GetPackageByID(id string) (*models.Package, error)
	GetPackagesByClassID(classID string) ([]models.Package, error)
	GetAllPackages(params dto.PackageQueryParam) ([]models.Package, int64, error)
	GetUserPackagesWithRemainingCredit(packageID string) ([]models.UserPackage, error)
}

type packageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(db *gorm.DB) PackageRepository {
	return &packageRepository{db}
}

func (r *packageRepository) DeletePackage(id string) error {
	return r.db.Delete(&models.Package{}, "id = ?", id).Error
}

func (r *packageRepository) CreatePackage(pkg *models.Package) error {
	return r.db.Create(pkg).Error
}

func (r *packageRepository) UpdatePackage(pkg *models.Package) error {
	return r.db.Save(pkg).Error
}

func (r *packageRepository) GetPackageByID(id string) (*models.Package, error) {
	var pkg models.Package
	err := r.db.Preload("Classes").First(&pkg, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &pkg, err
}

func (r *packageRepository) GetPackagesByClassID(classID string) ([]models.Package, error) {
	var packages []models.Package
	err := r.db.Joins("JOIN package_classes ON packages.id = package_classes.package_id").
		Where("package_classes.class_id = ?", classID).
		Preload("Classes").
		Find(&packages).Error
	return packages, err
}

func (r *packageRepository) GetAllPackages(params dto.PackageQueryParam) ([]models.Package, int64, error) {
	var packages []models.Package
	var count int64

	db := r.db.Model(&models.Package{}).Preload("Classes")

	// search by name or description
	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("packages.name LIKE ? OR packages.description LIKE ?", like, like)
	}

	// status filter
	if params.Status != "" && params.Status != "all" {
		if params.Status == "active" {
			db = db.Where("packages.is_active = ?", true)
		} else if params.Status == "inactive" {
			db = db.Where("packages.is_active = ?", false)
		}
	}

	// sort
	switch params.Sort {
	case "name_asc":
		db = db.Order("packages.name asc")
	case "name_desc":
		db = db.Order("packages.name desc")
	case "price_asc":
		db = db.Order("packages.price asc")
	case "price_desc":
		db = db.Order("packages.price desc")
	default:
		db = db.Order("packages.created_at desc")
	}

	// pagination
	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Limit(params.Limit).Offset(offset).Find(&packages).Error; err != nil {
		return nil, 0, err
	}

	return packages, count, nil
}

func (r *packageRepository) GetUserPackagesWithRemainingCredit(packageID string) ([]models.UserPackage, error) {
	var userPackages []models.UserPackage
	err := r.db.Where("package_id = ? AND remaining_credit > 0", packageID).Find(&userPackages).Error
	return userPackages, err
}
