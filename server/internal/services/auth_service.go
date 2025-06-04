package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
	"server/pkg/utils"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

type AuthService interface {
	SendOTP(email string) error
	Logout(refreshToken string) error
	Register(req *dto.RegisterRequest) error
	GoogleSignIn(idToken string) (*dto.AuthResponse, error)
	Login(req *dto.LoginRequest) (*dto.AuthResponse, error)
	VerifyOTP(email, otp string) (*dto.AuthResponse, error)
	GetUserProfile(userID string) (*dto.AuthMeResponse, error)
	RefreshToken(refreshToken string) (*dto.AuthResponse, error)
	GetGoogleOAuthURL() string
	HandleGoogleOAuthCallback(code string) (*dto.AuthResponse, error)
}

type authService struct {
	repo repositories.AuthRepository
	user repositories.UserRepository
}

func NewAuthService(repo repositories.AuthRepository, user repositories.UserRepository) AuthService {
	return &authService{repo, user}
}

func (s *authService) SendOTP(email string) error {
	_, err := config.RedisClient.Get(config.Ctx, "otp_data:"+email).Result()
	if err != nil {
		return customErr.NewNotFound("OTP data not found")
	}

	limitKey := "otp_resend_limit:" + email
	count, _ := config.RedisClient.Get(config.Ctx, limitKey).Int()
	if count >= 3 {
		return customErr.NewTooManyRequest("Too many OTP requests")
	}
	config.RedisClient.Incr(config.Ctx, limitKey)
	config.RedisClient.Expire(config.Ctx, limitKey, 30*time.Minute)

	otp := utils.GenerateOTP(6)
	if err := config.RedisClient.Set(config.Ctx, "otp:"+email, otp, 5*time.Minute).Err(); err != nil {
		return customErr.NewInternal("Failed to store OTP", err)
	}
	subject := "Your New OTP Code"
	body := fmt.Sprintf("Your new OTP is %s", otp)
	if err := utils.SendEmail(subject, email, otp, body); err != nil {
		return customErr.NewInternal("Failed to send OTP email", err)
	}

	return nil
}

func (s *authService) Logout(refreshToken string) error {
	return s.repo.DeleteAllUserRefreshTokens(refreshToken)
}

func (s *authService) VerifyOTP(email, otp string) (*dto.AuthResponse, error) {
	savedOtp, err := config.RedisClient.Get(config.Ctx, "otp:"+email).Result()
	if err != nil || savedOtp != otp {
		return nil, customErr.NewUnauthorized("OTP is invalid or has expired")
	}
	config.RedisClient.Del(config.Ctx, "otp:"+email)

	val, err := config.RedisClient.Get(config.Ctx, "otp_data:"+email).Result()
	if err != nil {
		return nil, customErr.NewUnauthorized("Session has expired")
	}
	var temp map[string]string
	if err := json.Unmarshal([]byte(val), &temp); err != nil {
		return nil, customErr.NewInternal("Failed to parse user data", err)
	}

	user := models.User{
		Email:    email,
		Fullname: temp["fullname"],
		Password: temp["password"],
		Avatar:   utils.RandomUserAvatar(temp["fullname"]),
	}
	if err := s.user.CreateUser(&user); err != nil {
		return nil, customErr.NewConflict("Email is already registered")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}

	session := &models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(session); err != nil {
		return nil, customErr.NewInternal("Failed to store refresh token", err)
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) GetUserProfile(userID string) (*dto.AuthMeResponse, error) {
	user, err := s.user.GetUserByID(userID)
	if err != nil {
		return nil, customErr.NewNotFound("User not found")
	}
	return &dto.AuthMeResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Fullname: user.Fullname,
		Avatar:   user.Avatar,
		Role:     user.Role,
	}, nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	redisKey := fmt.Sprintf("login:attempt:%s", req.Email)
	attempts, _ := config.RedisClient.Get(config.Ctx, redisKey).Int()
	if attempts >= 5 {
		return nil, customErr.NewTooManyRequest("Too many request, please try again in 30 minutes")
	}

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		config.RedisClient.Incr(config.Ctx, redisKey)
		config.RedisClient.Expire(config.Ctx, redisKey, 30*time.Minute)
		return nil, customErr.NewUnauthorized("Invalid email or password")
	}

	config.RedisClient.Del(config.Ctx, redisKey)

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}
	session := &models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(session); err != nil {
		return nil, customErr.NewInternal("Failed to store refresh token", err)
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err == nil {
		if user.Password == "-" {
			return customErr.NewAlreadyExist("Email already registered")
		}
		return customErr.NewAlreadyExist("Email already registered")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return customErr.ErrInternalServer
	}

	otp := utils.GenerateOTP(6)
	subject := "One-Time Password (OTP)"
	body := fmt.Sprintf("Your OTP code is %s", otp)
	if err := utils.SendEmail(subject, req.Email, otp, body); err != nil {
		return customErr.ErrInternalServer
	}
	if err := config.RedisClient.Set(config.Ctx, "otp:"+req.Email, otp, 5*time.Minute).Err(); err != nil {
		return customErr.ErrInternalServer
	}
	tempData := map[string]string{
		"fullname": req.Fullname,
		"password": hashedPassword,
		"email":    req.Email,
	}
	jsonStr, err := json.Marshal(tempData)
	if err != nil {
		return customErr.ErrInternalServer
	}
	if err := config.RedisClient.Set(config.Ctx, "otp_data:"+req.Email, jsonStr, 30*time.Minute).Err(); err != nil {
		return customErr.ErrInternalServer
	}
	return nil
}

func (s *authService) RefreshToken(refreshToken string) (*dto.AuthResponse, error) {
	_, err := utils.DecodeRefreshToken(refreshToken)
	if err != nil {
		return nil, customErr.ErrUnauthorized
	}
	tokenModel, err := s.repo.FindRefreshToken(refreshToken)
	if err != nil {
		return nil, customErr.NewNotFound("Refresh token not found")
	}
	if tokenModel.ExpiredAt.Before(time.Now()) {
		return nil, customErr.ErrUnauthorized
	}
	user, err := s.user.GetUserByID(tokenModel.UserID.String())
	if err != nil {
		return nil, customErr.NewNotFound("User not found")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}
	newRefreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}

	if err := s.repo.DeleteRefreshToken(refreshToken); err != nil {
		return nil, customErr.NewInternal("Failed to delete refresh token", err)
	}
	session := &models.Token{
		UserID:    user.ID,
		Token:     newRefreshToken,
		ExpiredAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(session); err != nil {
		return nil, customErr.NewInternal("Failed to store new refresh token", err)
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *authService) GoogleSignIn(idToken string) (*dto.AuthResponse, error) {
	payload, err := idtoken.Validate(context.Background(), idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return nil, customErr.ErrUnauthorized
	}
	email, ok := payload.Claims["email"].(string)
	if !ok || email == "" {
		return nil, customErr.NewBadRequest("Invalid email in ID token")
	}
	name, _ := payload.Claims["name"].(string)

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		user = &models.User{
			Email:    email,
			Password: "-",
			Role:     "customer",
			Fullname: name,
			Avatar:   utils.RandomUserAvatar(name),
		}
		if err := s.user.CreateUser(user); err != nil {
			return nil, customErr.NewInternal("Failed to create user", err)
		}
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}

	session := &models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(session); err != nil {
		return nil, customErr.NewInternal("Failed to store refresh token", err)
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) GetGoogleOAuthURL() string {
	return config.GoogleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (s *authService) HandleGoogleOAuthCallback(code string) (*dto.AuthResponse, error) {
	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, customErr.ErrUnauthorized
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, customErr.ErrUnauthorized
	}

	return s.GoogleSignIn(rawIDToken)
}
