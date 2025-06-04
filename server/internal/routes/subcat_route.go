package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SubcategoryRoutes(r *gin.Engine, h *handlers.SubcategoryHandler) {
	subcategory := r.Group("/api/subcategories")

	subcategory.GET("", h.GetAllSubcategories)
	subcategory.GET("/:id", h.GetSubcategoryByID)
	subcategory.GET("/category/:categoryId", h.GetSubcategoriesByCategoryID)

	admin := subcategory.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateSubcategory)
	admin.PUT("/:id", h.UpdateSubcategory)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteSubcategory)
}
