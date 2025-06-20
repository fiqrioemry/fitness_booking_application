// internal/routes/user_route.go
package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, h *handlers.UserHandler) {
	user := r.Group("users")
	// customer&instructor-endpoints
	user.Use(middleware.AuthRequired(), middleware.RoleOnly("customer", "instructor"))
	user.GET("/me", h.GetProfile)
	user.PUT("/me", h.UpdateProfile)
	user.PUT("/me/avatar", h.UpdateAvatar)

	// admin-endpoints
	admin := r.Group("admin/users")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.GET("", h.GetAllUsers)
	admin.GET("/:id", h.GetUserDetail)
	admin.GET("/stats", h.GetUserStats)

}
