package services

import (
	"server/internal/dto"
	"server/internal/repositories"
)

type DashboardService interface {
	GetSummary() (*dto.DashboardSummaryResponse, error)
	GetRevenueStats(rangeType string) (*dto.RevenueStatsResponse, error)
}

type dashboardService struct {
	repo repositories.DashboardRepository
}

func NewDashboardService(repo repositories.DashboardRepository) DashboardService {
	return &dashboardService{repo}
}

func (s *dashboardService) GetSummary() (*dto.DashboardSummaryResponse, error) {
	users, _ := s.repo.CountUsers()
	revenue, _ := s.repo.SumRevenue()
	classes, _ := s.repo.CountClasses()
	bookings, _ := s.repo.CountBookings()
	payments, _ := s.repo.CountPayments()
	instructors, _ := s.repo.CountInstructors()
	packages, _ := s.repo.CountActivePackages()
	absent, _ := s.repo.CountAttendanceByStatus("absent")
	attended, _ := s.repo.CountAttendanceByStatus("attended")

	return &dto.DashboardSummaryResponse{
		TotalUsers:         users,
		TotalInstructors:   instructors,
		TotalClasses:       classes,
		TotalBookings:      bookings,
		TotalPayments:      payments,
		TotalRevenue:       revenue,
		ActivePackages:     packages,
		TotalAttendance:    attended + absent,
		AttendedAttendance: attended,
		AbsentAttendance:   absent,
	}, nil
}

func (s *dashboardService) GetRevenueStats(rangeType string) (*dto.RevenueStatsResponse, error) {
	stats, total, err := s.repo.GetRevenueStatsByRange(rangeType)
	if err != nil {
		return nil, err
	}

	return &dto.RevenueStatsResponse{
		Range:         rangeType,
		TotalRevenue:  total,
		RevenueSeries: stats,
	}, nil
}
