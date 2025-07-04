package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ScheduleTemplateHandler struct {
	service services.ScheduleTemplateService
}

func NewScheduleTemplateHandler(service services.ScheduleTemplateService) *ScheduleTemplateHandler {
	return &ScheduleTemplateHandler{service}
}

func (h *ScheduleTemplateHandler) GetAllTemplates(c *gin.Context) {
	templates, err := h.service.GetAllTemplates()
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, templates)
}

func (h *ScheduleTemplateHandler) UpdateScheduleTemplate(c *gin.Context) {
	templateID := c.Param("id")
	var req dto.UpdateScheduleTemplateRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.UpdateScheduleTemplate(templateID, req); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Template updated successfully"})
}

func (h *ScheduleTemplateHandler) DeleteTemplate(c *gin.Context) {
	templateID := c.Param("id")
	if err := h.service.DeleteTemplate(templateID); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Template deleted successfully"})
}

func (h *ScheduleTemplateHandler) RunScheduleTemplate(c *gin.Context) {
	templateID := c.Param("id")
	if err := h.service.RunTemplate(templateID); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Template  Activated successfully"})
}

func (h *ScheduleTemplateHandler) StopScheduleTemplate(c *gin.Context) {
	templateID := c.Param("id")

	if err := h.service.StopTemplate(templateID); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Template deactivated successfully"})
}
