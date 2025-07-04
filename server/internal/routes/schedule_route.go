package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(r *gin.RouterGroup, h *handlers.ClassScheduleHandler) {
	// public-endpoints
	r.GET("/schedules", h.GetAllClassSchedules)

	// customer-endpoints
	customer := r.Group("/schedules")
	customer.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	customer.GET("/status", h.GetSchedulesWithStatus)
	customer.GET("/:id", h.GetScheduleByID)

	// instructor-endpoints
	instructor := r.Group("/instructor/schedules")
	instructor.Use(middleware.AuthRequired(), middleware.RoleOnly("instructor"))
	instructor.GET("", h.GetInstructorSchedules)
	instructor.PATCH("/:id/open", h.OpenClassSchedule)
	instructor.GET("/:id/attendance", h.GetClassAttendances)

	// admin-endpoints
	admin := r.Group("/admin/schedules")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateClassSchedule)
	admin.POST("/recurring", h.CreateRecurringSchedule)
	admin.PUT("/:id", h.UpdateClassSchedule)
	admin.DELETE("/:id", h.DeleteClassSchedule)
}
