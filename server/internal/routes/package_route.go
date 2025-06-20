package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func PackageRoutes(r *gin.RouterGroup, h *handlers.PackageHandler) {
	// public-endpoints
	r.GET("/packages", h.GetAllPackages)
	r.GET("/packages/:id", h.GetPackageByID)

	// admin-endpoints
	admin := r.Group("/admin/packages")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreatePackage)
	admin.PUT("/:id", h.UpdatePackage)
	admin.DELETE("/:id", h.DeletePackage)
}
