package repositories

import (
	"errors"
	"server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherRepository interface {
	CreateNewVoucher(v *models.Voucher) error
	GetAllVouchers() ([]models.Voucher, error)
	DeleteVoucherByID(id string) error
	UpdateVoucher(v *models.Voucher) error
	GetByCode(code string) (*models.Voucher, error)
	GetVoucherByID(id string) (*models.Voucher, error)
	InsertUsedVoucher(userID, voucherID uuid.UUID) error
	CheckVoucherUsed(userID, voucherID uuid.UUID) (bool, error)
	GetValidVoucherByCode(code string) (*models.Voucher, error)
}

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db}
}

func (r *voucherRepository) DeleteVoucherByID(id string) error {
	return r.db.Delete(&models.Voucher{}, "id = ?", id).Error
}

func (r *voucherRepository) CreateNewVoucher(v *models.Voucher) error {
	return r.db.Create(v).Error
}

func (r *voucherRepository) UpdateVoucher(v *models.Voucher) error {
	return r.db.Save(v).Error
}

func (r *voucherRepository) GetAllVouchers() ([]models.Voucher, error) {
	var vouchers []models.Voucher
	err := r.db.Order("created_at desc").Find(&vouchers).Error
	return vouchers, err
}

func (r *voucherRepository) GetByCode(code string) (*models.Voucher, error) {
	var v models.Voucher
	err := r.db.Where("code = ?", code).First(&v).Error
	return &v, err
}

func (r *voucherRepository) GetValidVoucherByCode(code string) (*models.Voucher, error) {
	var voucher models.Voucher
	err := r.db.
		Where("code = ? AND expired_at > NOW() AND quota > 0", code).
		First(&voucher).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &voucher, err
}

func (r *voucherRepository) CheckVoucherUsed(userID, voucherID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.UsedVoucher{}).
		Where("user_id = ? AND voucher_id = ?", userID, voucherID).
		Count(&count).Error
	return count > 0, err
}

func (r *voucherRepository) InsertUsedVoucher(userID, voucherID uuid.UUID) error {
	return r.db.Create(&models.UsedVoucher{
		ID:        uuid.New(),
		UserID:    userID,
		VoucherID: voucherID,
		UsedAt:    time.Now(),
	}).Error
}

func (r *voucherRepository) GetVoucherByID(id string) (*models.Voucher, error) {
	var voucher models.Voucher
	err := r.db.First(&voucher, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &voucher, err
}
