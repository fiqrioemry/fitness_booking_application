package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func LevelRoutes(r *gin.RouterGroup, h *handlers.LevelHandler) {
	// public-endpoints
	r.GET("/api/v1/levels", h.GetAllLevels)
	r.GET("/api/v1/levels/:id", h.GetLevelByID)

	// admin-endpoints
	admin := r.Group("/api/v1/admin/levels")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLevel)
	admin.PUT("/:id", h.UpdateLevel)
	admin.DELETE("/:id", h.DeleteLevel)
}
