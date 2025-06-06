package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ClassRoutes(r *gin.Engine, h *handlers.ClassHandler) {
	// public-endpoints
	r.GET("/api/classes", h.GetAllClasses)
	r.GET("/api/classes/:id", h.GetClassByID)

	// admin-endpoints
	admin := r.Group("/api/admin/classes")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateClass)
	admin.PUT("/:id", h.UpdateClass)
	admin.POST("/:id/gallery", h.UploadClassGallery)
	admin.DELETE("/:id", h.DeleteClass)
}
