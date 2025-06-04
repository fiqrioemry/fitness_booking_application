package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InstructorRoutes(r *gin.Engine, handler *handlers.InstructorHandler) {
	// public endpoints
	instructor := r.Group("/api/instructors")
	instructor.GET("", handler.GetAllInstructors)
	instructor.GET("/:id", handler.GetInstructorByID)

	// admin-protected endpoints
	admin := instructor.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", handler.CreateInstructor)
	admin.PUT("/:id", handler.UpdateInstructor)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), handler.DeleteInstructor)
}
