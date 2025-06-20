package routes

import (
	"server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func DashboardRoutes(r *gin.RouterGroup, handler *handlers.DashboardHandler) {
	admin := r.Group("/api/v1/admin")
	admin.GET("/dashboard/summary", handler.GetSummary)
	admin.GET("/dashboard/revenue", handler.GetRevenueStats)
}
