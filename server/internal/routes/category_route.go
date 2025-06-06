package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, h *handlers.CategoryHandler) {
	// public-endpoints
	r.GET("/api/categories", h.GetAllCategories)
	r.GET("/api/categories/:id", h.GetCategoryByID)

	// admin-endpoints
	admin := r.Group("/api/admin/categories")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateCategory)
	admin.PUT("/:id", h.UpdateCategory)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteCategory)
}
