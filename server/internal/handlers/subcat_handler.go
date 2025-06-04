package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type SubcategoryHandler struct {
	subcategoryService services.SubcategoryService
}

func NewSubcategoryHandler(subcategoryService services.SubcategoryService) *SubcategoryHandler {
	return &SubcategoryHandler{subcategoryService}
}

func (h *SubcategoryHandler) CreateSubcategory(c *gin.Context) {
	var req dto.CreateSubcategoryRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.subcategoryService.CreateSubcategory(req); err != nil {
		utils.HandleServiceError(c, err, "Failed to create subcategory")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Subcategory created successfully"})
}

func (h *SubcategoryHandler) UpdateSubcategory(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateSubcategoryRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.subcategoryService.UpdateSubcategory(id, req); err != nil {
		utils.HandleServiceError(c, err, "Failed to update subcategory")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subcategory updated successfully"})
}

func (h *SubcategoryHandler) DeleteSubcategory(c *gin.Context) {
	id := c.Param("id")

	if err := h.subcategoryService.DeleteSubcategory(id); err != nil {
		utils.HandleServiceError(c, err, "Failed to delete subcategory")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subcategory deleted successfully"})
}

func (h *SubcategoryHandler) GetAllSubcategories(c *gin.Context) {
	subcategories, err := h.subcategoryService.GetAllSubcategories()
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch subcategories")
		return
	}

	if subcategories == nil {
		subcategories = []dto.SubcategoryResponse{}
	}

	c.JSON(http.StatusOK, subcategories)
}

func (h *SubcategoryHandler) GetSubcategoryByID(c *gin.Context) {
	id := c.Param("id")

	subcategory, err := h.subcategoryService.GetSubcategoryByID(id)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch subcategory")
		return
	}

	c.JSON(http.StatusOK, subcategory)
}

func (h *SubcategoryHandler) GetSubcategoriesByCategoryID(c *gin.Context) {
	categoryID := c.Param("categoryId")

	subcategories, err := h.subcategoryService.GetSubcategoriesByCategoryID(categoryID)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch subcategories by category")
		return
	}

	if subcategories == nil {
		subcategories = []dto.SubcategoryResponse{}
	}

	c.JSON(http.StatusOK, subcategories)
}
