package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup, h *handlers.CategoryHandler) {
	// public-endpoints
	r.GET("/categories", h.GetAllCategories)
	r.GET("/categories/:id", h.GetCategoryByID)

	// admin-endpoints
	admin := r.Group("/admin/categories")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateCategory)
	admin.PUT("/:id", h.UpdateCategory)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteCategory)
}
