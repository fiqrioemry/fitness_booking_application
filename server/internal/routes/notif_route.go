package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.Engine, h *handlers.NotificationHandler) {
	notif := r.Group("/api/notifications")
	notif.Use(middleware.AuthRequired())
	notif.GET("/settings", h.GetNotificationSettings)
	notif.PUT("/settings", h.UpdateNotificationSetting)
	notif.GET("", h.GetAllNotifications)
	notif.PATCH("/read", h.MarkAllNotificationsAsRead)
	notif.POST("/broadcast", middleware.RoleOnly("admin"), h.SendNewNotificatioon)
}
