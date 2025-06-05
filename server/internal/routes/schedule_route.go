package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ClassScheduleRoutes(r *gin.Engine, h *handlers.ClassScheduleHandler) {
	schedule := r.Group("/api/schedules")

	// public endpoints
	schedule.GET("", h.GetAllClassSchedules)

	// base authentication
	auth := schedule.Group("")
	auth.Use(middleware.AuthRequired())

	// instructor-protected endpoints
	instructor := auth.Group("/instructor")
	instructor.Use(middleware.RoleOnly("instructor"))
	instructor.GET("", h.GetInstructorSchedules)
	instructor.PATCH("/:id/open", h.OpenClassSchedule)
	instructor.GET("/:id/attendance", h.GetClassAttendances)

	// customer-protected endpoints
	customer := auth.Group("")
	customer.Use(middleware.RoleOnly("customer"))
	customer.GET("/status", h.GetSchedulesWithStatus)
	customer.GET("/:id", h.GetScheduleByID)

	// admin-protected endpoints
	admin := auth.Group("")
	admin.Use(middleware.RoleOnly("admin"))
	admin.POST("", h.CreateClassSchedule)
	admin.POST("/recurring", h.CreateRecurringSchedule)
	admin.PUT("/:id", h.UpdateClassSchedule)
	admin.DELETE("/:id", h.DeleteClassSchedule)
}
