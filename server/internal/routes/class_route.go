package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ClassRoutes(r *gin.Engine, handler *handlers.ClassHandler) {
	class := r.Group("/api/classes")

	// public endpoints
	class.GET("", handler.GetAllClasses)
	class.GET("/:id", handler.GetClassByID)

	// admin-protected endpoints
	admin := class.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", handler.CreateClass)
	admin.PUT("/:id", handler.UpdateClass)
	admin.POST("/:id/gallery", handler.UploadClassGallery)
	admin.DELETE("/:id", handler.DeleteClass)
}
