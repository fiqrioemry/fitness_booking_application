// internal/routes/auth_route.go
package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, h *handlers.AuthHandler) {
	auth := r.Group("/api/v1/auth")

	// public-endpoints
	auth.POST("/login", h.Login)
	auth.POST("/logout", h.Logout)
	auth.POST("/register", h.Register)
	auth.POST("/send-otp", h.ResendOTP)
	auth.POST("/verify-otp", h.VerifyOTP)
	auth.POST("/refresh-token", h.RefreshToken)
	auth.GET("/google", h.GoogleOAuthRedirect)
	auth.GET("/google/callback", h.GoogleOAuthCallback)

	// authenticated-endpoint
	auth.GET("/me", middleware.AuthRequired(), h.AuthMe)
}
