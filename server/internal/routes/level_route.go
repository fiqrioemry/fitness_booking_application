package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func LevelRoutes(r *gin.Engine, h *handlers.LevelHandler) {
	// public-endpoints
	r.GET("/api/levels", h.GetAllLevels)
	r.GET("/api/levels/:id", h.GetLevelByID)

	// admin-endpoints
	admin := r.Group("/api/admin/levels")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLevel)
	admin.PUT("/:id", h.UpdateLevel)
	admin.DELETE("/:id", h.DeleteLevel)
}
