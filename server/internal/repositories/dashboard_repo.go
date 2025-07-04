package repositories

import (
	"server/internal/dto"
	"server/internal/models"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	CountUsers() (int, error)
	CountInstructors() (int, error)
	CountClasses() (int, error)
	CountBookings() (int, error)
	CountPayments() (int, error)
	SumRevenue() (float64, error)
	CountActivePackages() (int, error)

	GetRevenueStatsByRange(rangeType string) ([]dto.RevenueStat, float64, error)
	CountAttendanceByStatus(status string) (int, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db}
}

func (r *dashboardRepository) CountUsers() (int, error) {
	var count int64
	err := r.db.Model(&models.User{}).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountInstructors() (int, error) {
	var count int64
	err := r.db.Model(&models.Instructor{}).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountClasses() (int, error) {
	var count int64
	err := r.db.Model(&models.Class{}).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountBookings() (int, error) {
	var count int64
	err := r.db.Model(&models.Booking{}).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountPayments() (int, error) {
	var count int64
	err := r.db.Model(&models.Payment{}).Where("status = ?", "success").Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) SumRevenue() (float64, error) {
	var total float64
	err := r.db.Model(&models.Payment{}).Where("status = ?", "success").Select("SUM(total)").Scan(&total).Error
	return total, err
}

func (r *dashboardRepository) CountActivePackages() (int, error) {
	var count int64
	err := r.db.Model(&models.Package{}).Where("is_active = true").Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountAttendanceByStatus(status string) (int, error) {
	var count int64
	err := r.db.Model(&models.Attendance{}).Where("status = ?", status).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) GetRevenueStatsByRange(status string) ([]dto.RevenueStat, float64, error) {
	var stats []dto.RevenueStat
	var total float64

	query := r.db.Model(&models.Payment{}).Where("status = ?", "success")

	selectClause := ""
	groupClause := ""
	orderClause := ""

	switch status {
	case "daily":
		selectClause = "DATE(paid_at) as date"
		groupClause = "DATE(paid_at)"
		orderClause = "DATE(paid_at)"
	case "monthly":
		selectClause = "DATE_FORMAT(paid_at, '%Y-%m') as date"
		groupClause = "DATE_FORMAT(paid_at, '%Y-%m')"
		orderClause = "DATE_FORMAT(paid_at, '%Y-%m')"
	case "yearly":
		selectClause = "YEAR(paid_at) as date"
		groupClause = "YEAR(paid_at)"
		orderClause = "YEAR(paid_at)"
	default:
		selectClause = "DATE(paid_at) as date"
		groupClause = "DATE(paid_at)"
		orderClause = "DATE(paid_at)"
	}

	err := query.
		Select(selectClause + ", SUM(total) as total").
		Group(groupClause).
		Order(orderClause + " ASC").
		Scan(&stats).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Select("SUM(total)").Scan(&total).Error
	return stats, total, err
}
