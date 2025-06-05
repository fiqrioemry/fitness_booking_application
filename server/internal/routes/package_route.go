package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func PackageRoutes(r *gin.Engine, h *handlers.PackageHandler) {
	pkg := r.Group("/api/packages")
	{
		// public endpoints
		pkg.GET("", h.GetAllPackages)
		pkg.GET("/:id", h.GetPackageByID)

		// admin-protected endpoints
		admin := pkg.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
		admin.POST("", h.CreatePackage)
		admin.PUT("/:id", h.UpdatePackage)
		admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeletePackage)
	}
}
