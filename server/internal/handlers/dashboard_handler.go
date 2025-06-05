package handlers

import (
	"net/http"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboardService services.DashboardService
}

func NewDashboardHandler(dashboardService services.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService}
}

func (h *DashboardHandler) GetSummary(c *gin.Context) {
	data, err := h.dashboardService.GetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch dashboard summary"})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (h *DashboardHandler) GetRevenueStats(c *gin.Context) {
	rangeType := c.DefaultQuery("range", "daily")

	result, err := h.dashboardService.GetRevenueStats(rangeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
