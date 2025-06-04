package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	ExpireOldPendingPayments() (int64, error)
	CreatePayment(payment *models.Payment) error
	UpdatePayment(payment *models.Payment) error
	GetPaymentByID(id string) (*models.Payment, error)
	GetPaymentByOrderID(orderID string) (*models.Payment, error)
	GetAllUserPayments(params dto.PaymentQueryParam) ([]models.Payment, int64, error)
	GetPaymentsByUserID(userID string, params dto.PaymentQueryParam) ([]models.Payment, int64, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetPaymentByID(id string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, "id = ?", id).Error
	return &payment, err
}

func (r *paymentRepository) GetPaymentByOrderID(orderID string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, "id = ?", orderID).Error
	return &payment, err
}

func (r *paymentRepository) UpdatePayment(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func applyPaymentFilters(db *gorm.DB, params dto.PaymentQueryParam) *gorm.DB {
	if params.Search != "" {
		like := "%" + params.Search + "%"
		db = db.Where("email LIKE ? OR fullname LIKE ?", like, like)
	}
	if params.Status != "" && params.Status != "all" {
		db = db.Where("status = ?", params.Status)
	}

	if params.StartDate != "" && params.EndDate != "" {
		db = db.Where("paid_at BETWEEN ? AND ?", params.StartDate, params.EndDate)
	}

	switch params.Sort {
	case "paid_at_asc":
		db = db.Order("paid_at asc")
	case "paid_at_desc":
		db = db.Order("paid_at desc")
	case "name_asc":
		db = db.Order("fullname asc")
	case "name_desc":
		db = db.Order("fullname desc")
	case "email_asc":
		db = db.Order("email asc")
	case "email_desc":
		db = db.Order("email desc")
	default:
		db = db.Order("paid_at desc")
	}
	return db
}

func (r *paymentRepository) GetAllUserPayments(params dto.PaymentQueryParam) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var count int64

	db := r.db.Model(&models.Payment{})

	db = applyPaymentFilters(db, params)

	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Limit(params.Limit).Offset(offset).Find(&payments).Error; err != nil {
		return nil, 0, err
	}

	return payments, count, nil
}

func (r *paymentRepository) GetPaymentsByUserID(userID string, params dto.PaymentQueryParam) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var count int64

	db := r.db.Model(&models.Payment{})
	db = applyPaymentFilters(db, params)
	db = db.Where("user_id = ?", userID)

	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Limit(params.Limit).Offset(offset).Find(&payments).Error; err != nil {
		return nil, 0, err
	}

	return payments, count, nil
}

func (r *paymentRepository) ExpireOldPendingPayments() (int64, error) {
	threshold := time.Now().Add(-24 * time.Hour)

	result := r.db.Model(&models.Payment{}).
		Where("status = ? AND paid_at <= ?", "pending", threshold).
		Update("status", "failed")

	return result.RowsAffected, result.Error
}
