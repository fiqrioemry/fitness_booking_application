package routes

import (
	"server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func DashboardRoutes(r *gin.Engine, handler *handlers.DashboardHandler) {
	admin := r.Group("/api/admin")
	{
		admin.GET("/dashboard/summary", handler.GetSummary)
		admin.GET("/dashboard/revenue", handler.GetRevenueStats)
	}
}
