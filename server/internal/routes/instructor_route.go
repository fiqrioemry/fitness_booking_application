package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InstructorRoutes(r *gin.RouterGroup, h *handlers.InstructorHandler) {
	// public-endpoints
	r.GET("/api/v1/instructors", h.GetAllInstructors)
	r.GET("/api/v1/instructors/:id", h.GetInstructorByID)

	// admin-endpoints
	admin := r.Group("/api/v1/admin/instructors")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateInstructor)
	admin.PUT("/:id", h.UpdateInstructor)
	admin.DELETE("/:id", h.DeleteInstructor)
}
