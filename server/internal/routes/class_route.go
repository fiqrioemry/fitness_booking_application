package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ClassRoutes(r *gin.Engine, h *handlers.ClassHandler) {
	class := r.Group("/api/classes")

	// public endpoints
	class.GET("", h.GetAllClasses)
	class.GET("/:id", h.GetClassByID)

	// admin-protected endpoints
	admin := class.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateClass)
	admin.PUT("/:id", h.UpdateClass)
	admin.POST("/:id/gallery", h.UploadClassGallery)
	admin.DELETE("/:id", h.DeleteClass)
}
