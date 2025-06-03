// internal/routes/user_route.go
package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, handler *handlers.UserHandler) {
	user := r.Group("/api/users")
	user.Use(middleware.AuthRequired())
	user.GET("/me", handler.GetProfile)
	user.PUT("/me", handler.UpdateProfile)
	user.PUT("/me/avatar", handler.UpdateAvatar)

	admin := r.Group("/api/admin/users")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.GET("", handler.GetAllUsers)
	admin.GET("/:id", handler.GetUserDetail)

}
