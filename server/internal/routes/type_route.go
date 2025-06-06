// internal/routes/type_route.go
package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func TypeRoutes(r *gin.Engine, h *handlers.TypeHandler) {
	// public access
	t := r.Group("/api/types")
	t.GET("", h.GetAllTypes)
	t.GET("/:id", h.GetTypeByID)

	// admin-endpoints
	admin := r.Group("/api/admin/types")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateType)
	admin.PUT("/:id", h.UpdateType)
	admin.DELETE("/:id", h.DeleteType)
}
