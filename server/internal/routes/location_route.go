package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.Engine, h *handlers.LocationHandler) {
	locationGroup := r.Group("/api/locations")

	// public endpoints
	locationGroup.GET("", h.GetAllLocations)
	locationGroup.GET("/:id", h.GetLocationByID)

	// admin-protected endpoints
	admin := locationGroup.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLocation)
	admin.PUT("/:id", h.UpdateLocation)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteLocation)
}
