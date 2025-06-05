package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserPackageRoutes(r *gin.Engine, h *handlers.UserPackageHandler) {
	user := r.Group("/api/user-packages")
	user.Use(middleware.AuthRequired())
	user.GET("", h.GetUserPackages)
	user.GET("/class/:id", h.GetUserPackagesByClassID)
}
