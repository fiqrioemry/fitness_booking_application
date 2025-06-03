package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/repositories"
	"server/pkg/utils"
	"time"
)

type UserService interface {
	GetUserStats() (*dto.UserStatsResponse, error)
	GetUserDetail(id string) (*dto.UserDetailResponse, error)
	UpdateAvatar(userID string, req dto.UpdateAvatarRequest) error
	UpdateProfile(userID string, req dto.UpdateUserDetailRequest) error
	GetAllUsers(params dto.UserQueryParam) ([]dto.UserListResponse, *dto.PaginationResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUserStats() (*dto.UserStatsResponse, error) {
	total, customers, instructors, admins, newMonth, err := s.repo.GetUserStats()
	if err != nil {
		return nil, err
	}
	return &dto.UserStatsResponse{
		Total:        total,
		Customers:    customers,
		Instructors:  instructors,
		Admins:       admins,
		NewThisMonth: newMonth,
	}, nil
}

func (s *userService) UpdateAvatar(userID string, req dto.UpdateAvatarRequest) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Avatar != "" && user.Avatar != req.AvatarURL && !isDiceBear(user.Avatar) {
		_ = utils.DeleteFromCloudinary(user.Avatar)
	}

	user.Avatar = req.AvatarURL
	return s.repo.UpdateUser(user)
}

func (s *userService) UpdateProfile(userID string, req dto.UpdateUserDetailRequest) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}
	user.Bio = req.Bio
	user.Phone = req.Phone
	user.Gender = req.Gender
	user.Fullname = req.Fullname
	if req.Birthday != "" {
		birthday, err := time.Parse("2006-01-02", req.Birthday)
		if err == nil {
			user.Birthday = &birthday
		}
	}

	return s.repo.UpdateUser(user)
}

func (s *userService) GetAllUsers(params dto.UserQueryParam) ([]dto.UserListResponse, *dto.PaginationResponse, error) {

	users, total, err := s.repo.FindAllUsers(params)
	if err != nil {
		return nil, nil, err
	}

	var results []dto.UserListResponse
	for _, u := range users {
		results = append(results, dto.UserListResponse{
			ID:       u.ID.String(),
			Email:    u.Email,
			Role:     u.Role,
			Phone:    u.Phone,
			Avatar:   u.Avatar,
			Fullname: u.Fullname,
			JoinedAt: u.CreatedAt.Format("2006-01-02"),
		})
	}
	pagination := utils.Paginate(total, params.Page, params.Limit)
	return results, pagination, nil

}

func (s *userService) GetUserDetail(id string) (*dto.UserDetailResponse, error) {
	u, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	var lastLogin string
	if len(u.Tokens) > 0 {
		lastLogin = u.Tokens[len(u.Tokens)-1].CreatedAt.Format(time.RFC3339)
	}

	res := &dto.UserDetailResponse{
		ID:        u.ID.String(),
		Email:     u.Email,
		Role:      u.Role,
		Fullname:  u.Fullname,
		Phone:     u.Phone,
		Avatar:    u.Avatar,
		Gender:    u.Gender,
		Bio:       u.Bio,
		LastLogin: lastLogin,
		JoinedAt:  u.CreatedAt.Format("2006-01-02"),
	}

	if u.Birthday != nil {
		res.Birthday = u.Birthday.Format("2006-01-02")
	}

	return res, nil
}

func isDiceBear(url string) bool {
	return url != "" && (len(url) > 0 && (url[:30] == "https://api.dicebear.com" || url[:31] == "https://avatars.dicebear.com"))
}
