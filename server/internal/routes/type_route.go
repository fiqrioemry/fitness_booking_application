// internal/routes/type_route.go
package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func TypeRoutes(r *gin.RouterGroup, h *handlers.TypeHandler) {
	// public-endpoints
	t := r.Group("//types")
	t.GET("", h.GetAllTypes)
	t.GET("/:id", h.GetTypeByID)

	// admin-endpoints
	admin := r.Group("/admin/types")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateType)
	admin.PUT("/:id", h.UpdateType)
	admin.DELETE("/:id", h.DeleteType)
}
