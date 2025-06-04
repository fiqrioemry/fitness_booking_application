package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
	"server/pkg/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherService interface {
	DeleteVoucher(id string) error
	CreateVoucher(dto.CreateVoucherRequest) error
	GetAllVouchers() ([]dto.VoucherResponse, error)
	DecreaseQuota(userID uuid.UUID, code string) error
	UpdateVoucher(id string, req dto.UpdateVoucherRequest) error
	ApplyVoucher(req dto.ApplyVoucherRequest) (*dto.ApplyVoucherResponse, error)
}

type voucherService struct {
	repo repositories.VoucherRepository
}

func NewVoucherService(repo repositories.VoucherRepository) VoucherService {
	return &voucherService{repo}
}

func (s *voucherService) DeleteVoucher(id string) error {
	_, err := s.repo.GetVoucherByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return customErr.ErrNotFound
	}
	if err != nil {
		return customErr.NewInternal("failed to get voucher by ID", err)
	}

	if err := s.repo.DeleteVoucherByID(id); err != nil {
		return customErr.NewInternal("failed to delete voucher", err)
	}
	return nil
}

func (s *voucherService) CreateVoucher(req dto.CreateVoucherRequest) error {
	expiredAt, err := utils.ParseDate(req.ExpiredAt)
	if err != nil {
		return customErr.NewBadRequest("invalid date format")
	}

	voucher := models.Voucher{
		Code:         req.Code,
		Description:  req.Description,
		DiscountType: req.DiscountType,
		Discount:     req.Discount,
		MaxDiscount:  req.MaxDiscount,
		IsReusable:   req.IsReusable,
		Quota:        req.Quota,
		ExpiredAt:    expiredAt,
	}

	if err := s.repo.CreateNewVoucher(&voucher); err != nil {
		return customErr.NewInternal("failed to create voucher", err)
	}
	return nil
}

func (s *voucherService) GetAllVouchers() ([]dto.VoucherResponse, error) {
	vouchers, err := s.repo.GetAllVouchers()
	if err != nil {
		return nil, customErr.NewInternal("failed to fetch vouchers", err)
	}

	var result []dto.VoucherResponse
	for _, v := range vouchers {
		result = append(result, dto.VoucherResponse{
			ID:           v.ID.String(),
			Code:         v.Code,
			Description:  v.Description,
			DiscountType: v.DiscountType,
			Discount:     v.Discount,
			IsReusable:   v.IsReusable,
			MaxDiscount:  v.MaxDiscount,
			Quota:        v.Quota,
			ExpiredAt:    v.ExpiredAt.Format("2006-01-02"),
			CreatedAt:    v.CreatedAt.Format("2006-01-02"),
		})
	}

	return result, nil
}

func (s *voucherService) DecreaseQuota(userID uuid.UUID, code string) error {
	voucher, err := s.repo.GetByCode(code)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return customErr.ErrNotFound
	}
	if err != nil {
		return customErr.NewInternal("failed to get voucher by code", err)
	}

	if !voucher.IsReusable {
		if err := s.repo.InsertUsedVoucher(userID, voucher.ID); err != nil {
			return customErr.NewInternal("failed to insert used voucher", err)
		}
	}

	if voucher.Quota > 0 {
		voucher.Quota--
		if err := s.repo.UpdateVoucher(voucher); err != nil {
			return customErr.NewInternal("failed to update voucher quota", err)
		}
	}
	return nil
}

func (s *voucherService) UpdateVoucher(id string, req dto.UpdateVoucherRequest) error {
	expiredAt, err := utils.ParseDate(req.ExpiredAt)
	if err != nil {
		return customErr.NewBadRequest("invalid date format")
	}

	voucher, err := s.repo.GetVoucherByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return customErr.ErrNotFound
	}
	if err != nil {
		return customErr.NewInternal("failed to get voucher", err)
	}

	voucher.Quota = req.Quota
	voucher.Discount = req.Discount
	voucher.ExpiredAt = expiredAt
	voucher.IsReusable = req.IsReusable
	voucher.MaxDiscount = req.MaxDiscount
	voucher.Description = req.Description
	voucher.DiscountType = req.DiscountType

	if err := s.repo.UpdateVoucher(voucher); err != nil {
		return customErr.NewInternal("failed to update voucher", err)
	}
	return nil
}

func (s *voucherService) ApplyVoucher(req dto.ApplyVoucherRequest) (*dto.ApplyVoucherResponse, error) {
	voucher, err := s.repo.GetValidVoucherByCode(req.Code)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, customErr.NewBadRequest("invalid or expired voucher")
	}
	if err != nil {
		return nil, customErr.NewInternal("failed to fetch voucher", err)
	}

	var userUUID uuid.UUID
	hasUser := req.UserID != nil && *req.UserID != ""

	if hasUser {
		userUUID, err = uuid.Parse(*req.UserID)
		if err != nil {
			return nil, customErr.NewBadRequest("invalid user id")
		}

		if !voucher.IsReusable {
			used, err := s.repo.CheckVoucherUsed(userUUID, voucher.ID)
			if err != nil {
				return nil, customErr.NewInternal("failed to check voucher usage", err)
			}
			if used {
				return nil, customErr.NewConflict("voucher already used")
			}
		}
	}

	var discountValue float64
	if voucher.DiscountType == "percentage" {
		discountValue = req.Total * (voucher.Discount / 100)
		if voucher.MaxDiscount != nil && discountValue > *voucher.MaxDiscount {
			discountValue = *voucher.MaxDiscount
		}
	} else {
		discountValue = min(voucher.Discount, req.Total)
	}

	final := req.Total - discountValue

	return &dto.ApplyVoucherResponse{
		Code:          voucher.Code,
		DiscountType:  voucher.DiscountType,
		Discount:      voucher.Discount,
		MaxDiscount:   voucher.MaxDiscount,
		DiscountValue: discountValue,
		FinalTotal:    final,
	}, nil
}
