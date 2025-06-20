package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserPackageRoutes(r *gin.RouterGroup, h *handlers.UserPackageHandler) {
	user := r.Group("/api/v1/user-packages")
	// customer-endpoints
	user.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	user.GET("", h.GetUserPackages)
	user.GET("/class/:id", h.GetUserPackagesByClassID)
}
