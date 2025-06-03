package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/pkg/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
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

	//
	GetGoogleOAuthURL() string
	HandleGoogleOAuthCallback(code string) (*dto.AuthResponse, error)
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) SendOTP(email string) error {
	_, err := config.RedisClient.Get(config.Ctx, "otp_data:"+email).Result()
	if err != nil {
		return errors.New("no pending registration found for this email")
	}

	// Limit max. 3 times within 30 min.
	limitKey := "otp_resend_limit:" + email
	count, _ := config.RedisClient.Get(config.Ctx, limitKey).Int()
	if count >= 3 {
		return errors.New("OTP resend limit reached. Please wait")
	}
	config.RedisClient.Incr(config.Ctx, limitKey)
	config.RedisClient.Expire(config.Ctx, limitKey, 30*time.Minute)

	// Generate OTP and save
	otp := utils.GenerateOTP(6)
	if err := config.RedisClient.Set(config.Ctx, "otp:"+email, otp, 5*time.Minute).Err(); err != nil {
		return err
	}
	// send otp via email
	subject := "Your New OTP Code"
	body := fmt.Sprintf("Your new OTP is %s", otp)
	if err := utils.SendEmail(subject, email, otp, body); err != nil {
		return errors.New("failed to send OTP")
	}

	return nil
}

func (s *authService) Logout(refreshToken string) error {
	return s.repo.DeleteAllUserRefreshTokens(refreshToken)
}

func (s *authService) VerifyOTP(email, otp string) (*dto.AuthResponse, error) {
	// validate otp
	savedOtp, err := config.RedisClient.Get(config.Ctx, "otp:"+email).Result()
	if err != nil || savedOtp != otp {
		return nil, errors.New("invalid or expired OTP")
	}
	config.RedisClient.Del(config.Ctx, "otp:"+email)

	val, err := config.RedisClient.Get(config.Ctx, "otp_data:"+email).Result()
	if err != nil {
		return nil, errors.New("registration data expired or not found")
	}
	var temp map[string]string
	if err := json.Unmarshal([]byte(val), &temp); err != nil {
		return nil, errors.New("failed to parse user data")
	}
	fullname := temp["fullname"]
	hashedPassword := temp["password"]

	user := models.User{
		Email:    email,
		Fullname: fullname,
		Password: hashedPassword,
		Avatar:   utils.RandomUserAvatar(fullname),
	}
	if err := s.repo.CreateUser(&user); err != nil {
		return nil, err
	}
	// Generate token
	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	// save sessions
	session := &models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(session); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) GetUserProfile(userID string) (*dto.AuthMeResponse, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	response := dto.AuthMeResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Fullname: user.Fullname,
		Avatar:   user.Avatar,
		Role:     user.Role,
	}

	return &response, nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	redisKey := fmt.Sprintf("login:attempt:%s", req.Email)
	attemptsStr, _ := config.RedisClient.Get(config.Ctx, redisKey).Result()
	attempts := 0
	if attemptsStr != "" {
		attempts, _ = strconv.Atoi(attemptsStr)
	}

	if attempts >= 5 {
		return nil, errors.New("too many failed attempts, please try again in 30 minutes")
	}

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		config.RedisClient.Incr(config.Ctx, redisKey)
		config.RedisClient.Expire(config.Ctx, redisKey, 30*time.Minute)
		return nil, errors.New("invalid email or password")
	}

	config.RedisClient.Del(config.Ctx, redisKey)

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	tokenModel := &models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().UTC().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(tokenModel); err != nil {
		return nil, err
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
			return errors.New("email already registered via Google Sign-In")
		}
		return errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	otp := utils.GenerateOTP(6)
	subject := "One-Time Password (OTP)"
	body := fmt.Sprintf("Your OTP code is %s", otp)

	if err := utils.SendEmail(subject, req.Email, otp, body); err != nil {
		return errors.New("failed to send email")
	}

	if err := config.RedisClient.Set(config.Ctx, "otp:"+req.Email, otp, 5*time.Minute).Err(); err != nil {
		return err
	}

	tempData := map[string]string{
		"fullname": req.Fullname,
		"password": hashedPassword,
		"email":    req.Email,
	}

	jsonStr, err := json.Marshal(tempData)
	if err != nil {
		return err
	}

	if err := config.RedisClient.Set(config.Ctx, "otp_data:"+req.Email, jsonStr, 30*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}

func (s *authService) RefreshToken(refreshToken string) (*dto.AuthResponse, error) {

	_, err := utils.DecodeRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	tokenModel, err := s.repo.FindRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("refresh token not found")
	}

	if tokenModel.ExpiredAt.Before(time.Now().UTC()) {
		return nil, errors.New("refresh token expired")
	}

	user, err := s.repo.GetUserByID(tokenModel.UserID.String())
	if err != nil {
		return nil, errors.New("user not found")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	if err := s.repo.DeleteRefreshToken(refreshToken); err != nil {
		return nil, err
	}

	newToken := &models.Token{
		UserID:    user.ID,
		Token:     newRefreshToken,
		ExpiredAt: time.Now().UTC().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(newToken); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *authService) GoogleSignIn(idToken string) (*dto.AuthResponse, error) {
	payload, err := idtoken.Validate(context.Background(), idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return nil, errors.New("invalid Google ID token")
	}

	email, ok := payload.Claims["email"].(string)
	if !ok || email == "" {
		return nil, errors.New("email not found in token")
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

		if err := s.repo.CreateUser(user); err != nil {
			return nil, err
		}

		if user.ID == uuid.Nil {
			return nil, errors.New("failed to assign UUID to user")
		}

	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	tokenModel := &models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().UTC().Add(7 * 24 * time.Hour),
	}

	if tokenModel.UserID == uuid.Nil {
		return nil, errors.New("user ID is empty")
	}

	if err := s.repo.StoreRefreshToken(tokenModel); err != nil {
		return nil, err
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
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("missing id_token in Google response")
	}

	return s.GoogleSignIn(rawIDToken)
}
