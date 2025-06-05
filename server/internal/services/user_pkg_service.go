package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/pkg/utils"
	"time"
)

type UserPackageService interface {
	GetUserPackagesByClassID(userID, classID string) ([]dto.UserPackageResponse, error)
	GetUserPackages(userID string, params dto.PackageQueryParam) ([]dto.UserPackageResponse, *dto.PaginationResponse, error)
}

type userPackageService struct {
	repo repositories.UserPackageRepository
}

func NewUserPackageService(repo repositories.UserPackageRepository) UserPackageService {
	return &userPackageService{repo}
}

func (s *userPackageService) GetUserPackages(userID string, params dto.PackageQueryParam) ([]dto.UserPackageResponse, *dto.PaginationResponse, error) {

	pkgs, total, err := s.repo.GetUserPackages(userID, params)
	if err != nil {
		return nil, nil, err
	}

	if len(pkgs) == 0 {
		pkgs = make([]models.UserPackage, 0)
	}

	var results []dto.UserPackageResponse
	for _, p := range pkgs {
		expired := ""
		if p.ExpiredAt != nil {
			expired = p.ExpiredAt.Format("2006-01-02")
		}
		results = append(results, dto.UserPackageResponse{
			ID:              p.ID.String(),
			PackageName:     p.PackageName,
			RemainingCredit: p.RemainingCredit,
			ExpiredAt:       expired,
			PurchasedAt:     p.PurchasedAt.Format("2006-01-02"),
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return results, pagination, nil
}

func (s *userPackageService) GetUserPackagesByClassID(userID, classID string) ([]dto.UserPackageResponse, error) {
	userPackages, err := s.repo.GetUserPackagesByClassID(userID, classID)
	if err != nil {
		return nil, err
	}

	var result []dto.UserPackageResponse
	for _, up := range userPackages {
		var (
			expiredAt     string
			expiredInDays int
		)

		if up.ExpiredAt != nil {
			expiredAt = up.ExpiredAt.Format("2006-01-02")
			expiredInDays = int(max(0, int(time.Until(*up.ExpiredAt).Hours()/24)))
		}

		result = append(result, dto.UserPackageResponse{
			ID:              up.ID.String(),
			PackageID:       up.Package.ID.String(),
			PackageName:     up.Package.Name,
			RemainingCredit: up.RemainingCredit,
			ExpiredAt:       expiredAt,
			ExpiredInDays:   expiredInDays,
			PurchasedAt:     up.PurchasedAt.Format("2006-01-02"),
		})
	}

	return result, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
