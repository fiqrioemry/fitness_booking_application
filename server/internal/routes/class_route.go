package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ClassRoutes(r *gin.RouterGroup, h *handlers.ClassHandler) {
	// public-endpoints
	r.GET("/classes", h.GetAllClasses)
	r.GET("/classes/:id", h.GetClassByID)

	// admin-endpoints
	admin := r.Group("/admin/classes")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateClass)
	admin.PUT("/:id", h.UpdateClass)
	admin.POST("/:id/gallery", h.UploadClassGallery)
	admin.DELETE("/:id", h.DeleteClass)
}
