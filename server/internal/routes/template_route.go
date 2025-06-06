package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func TemplateRoutes(r *gin.Engine, handler *handlers.ScheduleTemplateHandler) {
	admin := r.Group("/api/admin/schedule-templates")
	// admin-endpoints
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.GET("", handler.GetAllTemplates)
	admin.PUT("/:id", handler.UpdateScheduleTemplate)
	admin.POST("/:id/run", handler.RunScheduleTemplate)
	admin.POST("/:id/stop", handler.StopScheduleTemplate)
	admin.DELETE("/:id", handler.DeleteTemplate)
}
