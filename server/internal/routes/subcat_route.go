// internal/routes/subcategory_route.go
package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SubcategoryRoutes(r *gin.RouterGroup, h *handlers.SubcategoryHandler) {
	// public-endpoints
	s := r.Group("/subcategories")
	s.GET("", h.GetAllSubcategories)
	s.GET("/:id", h.GetSubcategoryByID)
	s.GET("/category/:categoryId", h.GetSubcategoriesByCategoryID)

	// admin-endpoints
	admin := r.Group("/admin/subcategories")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateSubcategory)
	admin.PUT("/:id", h.UpdateSubcategory)
	admin.DELETE("/:id", h.DeleteSubcategory)
}
