package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TypeHandler struct {
	typeService services.TypeService
}

func NewTypeHandler(typeService services.TypeService) *TypeHandler {
	return &TypeHandler{typeService}
}

func (h *TypeHandler) CreateType(c *gin.Context) {
	var req dto.CreateTypeRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.typeService.CreateType(req); err != nil {
		utils.HandleServiceError(c, err, "Failed to create type")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Type created successfully"})
}

func (h *TypeHandler) UpdateType(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateTypeRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.typeService.UpdateType(id, req); err != nil {
		utils.HandleServiceError(c, err, "Failed to update type")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Type updated successfully"})
}

func (h *TypeHandler) DeleteType(c *gin.Context) {
	id := c.Param("id")

	if err := h.typeService.DeleteType(id); err != nil {
		utils.HandleServiceError(c, err, "Failed to delete type")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Type deleted successfully"})
}

func (h *TypeHandler) GetAllTypes(c *gin.Context) {
	types, err := h.typeService.GetAllTypes()
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch types")
		return
	}

	if types == nil {
		types = []dto.TypeResponse{}
	}

	c.JSON(http.StatusOK, types)
}

func (h *TypeHandler) GetTypeByID(c *gin.Context) {
	id := c.Param("id")

	typeClass, err := h.typeService.GetTypeByID(id)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch type")
		return
	}

	c.JSON(http.StatusOK, typeClass)
}
