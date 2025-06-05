package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InstructorRoutes(r *gin.Engine, h *handlers.InstructorHandler) {
	// public endpoints
	instructor := r.Group("/api/instructors")
	instructor.GET("", h.GetAllInstructors)
	instructor.GET("/:id", h.GetInstructorByID)

	// admin-protected endpoints
	admin := instructor.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateInstructor)
	admin.PUT("/:id", h.UpdateInstructor)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteInstructor)
}
