package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func LevelRoutes(r *gin.Engine, h *handlers.LevelHandler) {
	// public endpoints
	levelGroup := r.Group("/api/levels")
	levelGroup.GET("", h.GetAllLevels)
	levelGroup.GET("/:id", h.GetLevelByID)

	// admin-protected endpoints
	admin := levelGroup.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLevel)
	admin.PUT("/:id", h.UpdateLevel)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteLevel)
}
