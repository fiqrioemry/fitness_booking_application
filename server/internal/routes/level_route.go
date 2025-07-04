package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func LevelRoutes(r *gin.RouterGroup, h *handlers.LevelHandler) {
	// public-endpoints
	r.GET("/levels", h.GetAllLevels)
	r.GET("/levels/:id", h.GetLevelByID)

	// admin-endpoints
	admin := r.Group("/admin/levels")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLevel)
	admin.PUT("/:id", h.UpdateLevel)
	admin.DELETE("/:id", h.DeleteLevel)
}
