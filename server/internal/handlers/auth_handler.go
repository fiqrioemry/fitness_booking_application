package handlers

import (
	"net/http"
	"os"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) ResendOTP(c *gin.Context) {
	var req dto.SendOTPRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.SendOTP(req.Email); err != nil {
		utils.HandleServiceError(c, err, "Failed to resend OTP")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to email"})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.Register(&req); err != nil {
		utils.HandleServiceError(c, err, "Registration failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to your email"})
}

func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	tokens, err := h.service.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		utils.HandleServiceError(c, err, "OTP verification failed")
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)
	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	tokens, err := h.service.Login(&req)
	if err != nil {
		utils.HandleServiceError(c, err, "Login failed")
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)
	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		utils.HandleServiceError(c, err, "Refresh token missing")
		return
	}

	if err := h.service.Logout(refreshToken); err != nil {
		utils.HandleServiceError(c, err, "Logout failed")
		return
	}

	utils.ClearAccessTokenCookie(c)
	utils.ClearRefreshTokenCookie(c)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		utils.HandleServiceError(c, err, "Refresh token missing")
		return
	}

	tokens, err := h.service.RefreshToken(refreshToken)
	if err != nil {
		utils.HandleServiceError(c, err, "Token refresh failed")
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)
	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}

func (h *AuthHandler) AuthMe(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	response, err := h.service.GetUserProfile(userID)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to retrieve user profile")
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) GoogleOAuthRedirect(c *gin.Context) {
	url := h.service.GetGoogleOAuthURL()
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) GoogleOAuthCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authorization code is missing"})
		return
	}

	tokens, err := h.service.HandleGoogleOAuthCallback(code)
	if err != nil {
		utils.HandleServiceError(c, err, "Google OAuth failed")
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)
	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_REDIRECT_URL"))
}
