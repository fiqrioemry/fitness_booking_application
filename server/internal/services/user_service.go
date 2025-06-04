package services

import (
	"server/internal/dto"
	"server/internal/repositories"
	customErr "server/pkg/errors"
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
		return nil, customErr.NewInternal("failed to fetch user statistics", err)
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
		return customErr.NewNotFound("user not found")
	}

	if user.Avatar != "" && user.Avatar != req.AvatarURL && !isDiceBear(user.Avatar) {
		_ = utils.DeleteFromCloudinary(user.Avatar) // ignore error
	}

	user.Avatar = req.AvatarURL
	if err := s.repo.UpdateUser(user); err != nil {
		return customErr.NewInternal("failed to update avatar", err)
	}
	return nil
}

func (s *userService) UpdateProfile(userID string, req dto.UpdateUserDetailRequest) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return customErr.NewNotFound("user not found")
	}

	user.Bio = req.Bio
	user.Phone = req.Phone
	user.Gender = req.Gender
	user.Fullname = req.Fullname

	if req.Birthday != "" {
		birthday, err := time.Parse("2006-01-02", req.Birthday)
		if err != nil {
			return customErr.NewBadRequest("invalid birthday format")
		}
		user.Birthday = &birthday
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return customErr.NewInternal("failed to update user profile", err)
	}
	return nil
}

func (s *userService) GetAllUsers(params dto.UserQueryParam) ([]dto.UserListResponse, *dto.PaginationResponse, error) {
	users, total, err := s.repo.FindAllUsers(params)
	if err != nil {
		return nil, nil, customErr.NewInternal("failed to fetch user list", err)
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
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, customErr.NewNotFound("user not found")
	}

	var lastLogin string
	if len(user.Tokens) > 0 {
		lastLogin = user.Tokens[len(user.Tokens)-1].CreatedAt.Format(time.RFC3339)
	}

	res := &dto.UserDetailResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Role:      user.Role,
		Fullname:  user.Fullname,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		Bio:       user.Bio,
		LastLogin: lastLogin,
		JoinedAt:  user.CreatedAt.Format("2006-01-02"),
	}

	if user.Birthday != nil {
		res.Birthday = user.Birthday.Format("2006-01-02")
	}

	return res, nil
}

func isDiceBear(url string) bool {
	return url != "" && (len(url) > 30 && (url[:30] == "https://api.dicebear.com" || url[:31] == "https://avatars.dicebear.com"))
}
