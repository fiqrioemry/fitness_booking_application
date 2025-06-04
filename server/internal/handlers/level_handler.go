package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LevelHandler struct {
	levelService services.LevelService
}

func NewLevelHandler(levelService services.LevelService) *LevelHandler {
	return &LevelHandler{levelService}
}

func (h *LevelHandler) CreateLevel(c *gin.Context) {
	var req dto.CreateLevelRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.levelService.CreateLevel(req); err != nil {
		utils.HandleServiceError(c, err, "Failed to create level")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Level created successfully"})
}

func (h *LevelHandler) UpdateLevel(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateLevelRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.levelService.UpdateLevel(id, req); err != nil {
		utils.HandleServiceError(c, err, "Failed to update level")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Level updated successfully"})
}

func (h *LevelHandler) DeleteLevel(c *gin.Context) {
	id := c.Param("id")

	if err := h.levelService.DeleteLevel(id); err != nil {
		utils.HandleServiceError(c, err, "Failed to delete level")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Level deleted successfully"})
}

func (h *LevelHandler) GetAllLevels(c *gin.Context) {
	levels, err := h.levelService.GetAllLevels()
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch levels")
		return
	}

	if levels == nil {
		levels = []dto.LevelResponse{}
	}

	c.JSON(http.StatusOK, levels)
}

func (h *LevelHandler) GetLevelByID(c *gin.Context) {
	id := c.Param("id")

	level, err := h.levelService.GetLevelByID(id)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch level")
		return
	}

	c.JSON(http.StatusOK, level)
}
