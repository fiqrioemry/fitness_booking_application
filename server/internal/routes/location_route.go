package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.RouterGroup, h *handlers.LocationHandler) {
	// public-endpoints
	r.GET("/api/v1/locations", h.GetAllLocations)
	r.GET("/api/v1/locations/:id", h.GetLocationByID)

	// admin-endpoints
	admin := r.Group("/api/v1/admin/locations")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLocation)
	admin.PUT("/:id", h.UpdateLocation)
	admin.DELETE("/:id", h.DeleteLocation)
}
