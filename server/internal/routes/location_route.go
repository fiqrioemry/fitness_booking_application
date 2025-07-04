package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.RouterGroup, h *handlers.LocationHandler) {
	// public-endpoints
	r.GET("/locations", h.GetAllLocations)
	r.GET("/locations/:id", h.GetLocationByID)

	// admin-endpoints
	admin := r.Group("/admin/locations")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLocation)
	admin.PUT("/:id", h.UpdateLocation)
	admin.DELETE("/:id", h.DeleteLocation)
}
