package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func TemplateRoutes(r *gin.Engine, handler *handlers.ScheduleTemplateHandler) {
	template := r.Group("/api/schedule-templates")
	template.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))

	template.GET("", handler.GetAllTemplates)
	template.PUT("/:id", handler.UpdateScheduleTemplate)
	template.POST("/:id/run", handler.RunScheduleTemplate)
	template.POST("/:id/stop", handler.StopScheduleTemplate)
	template.DELETE("/:id", middleware.RoleOnly("owner"), handler.DeleteTemplate)
}
