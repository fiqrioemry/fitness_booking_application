package routes

import (
	"server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, handler *handlers.AuthHandler) {
	auth := r.Group("/api/auth")
	{
		// public endpoints
		auth.POST("/send-otp", handler.ResendOTP)
		auth.POST("/verify-otp", handler.VerifyOTP)
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.POST("/refresh-token", handler.RefreshToken)
		auth.GET("/google", handler.GoogleOAuthRedirect)
		auth.GET("/google/callback", handler.GoogleOAuthCallback)
	}
}
