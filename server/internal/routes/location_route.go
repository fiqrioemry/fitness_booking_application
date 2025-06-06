package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.Engine, h *handlers.LocationHandler) {
	// public-endpoints
	r.GET("/api/locations", h.GetAllLocations)
	r.GET("/api/locations/:id", h.GetLocationByID)

	// admin-endpoints
	admin := r.Group("/api/admin/locations")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLocation)
	admin.PUT("/:id", h.UpdateLocation)
	admin.DELETE("/:id", h.DeleteLocation)
}
