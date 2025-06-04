package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func PackageRoutes(r *gin.Engine, handler *handlers.PackageHandler) {
	pkg := r.Group("/api/packages")
	{
		// public endpoints
		pkg.GET("", handler.GetAllPackages)
		pkg.GET("/:id", handler.GetPackageByID)

		// admin-protected endpoints
		admin := pkg.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
		admin.POST("", handler.CreatePackage)
		admin.PUT("/:id", handler.UpdatePackage)
		admin.DELETE("/:id", middleware.RoleOnly("owner"), handler.DeletePackage)
	}
}
