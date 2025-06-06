package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.Engine, h *handlers.NotificationHandler) {
	// customer-endpoints
	notif := r.Group("/api/notifications")
	notif.Use(middleware.AuthRequired(), middleware.RoleOnly("customer", "instructor"))
	notif.GET("/settings", h.GetNotificationSettings)
	notif.PUT("/settings", h.UpdateNotificationSetting)
	notif.GET("", h.GetAllNotifications)
	notif.PATCH("/read", h.MarkAllNotificationsAsRead)

	// admin-endpoints
	admin := r.Group("/api/admin/notifications")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("/broadcast", h.SendNewNotificatioon)
}
