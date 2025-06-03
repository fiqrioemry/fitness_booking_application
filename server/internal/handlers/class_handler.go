package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClassHandler struct {
	service services.ClassService
}

func NewClassHandler(service services.ClassService) *ClassHandler {
	return &ClassHandler{service}
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var req dto.CreateClassRequest

	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	req.IsActive, _ = utils.ParseBoolFormField(c, "isActive")

	imageURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	req.ImageURL = imageURL

	if len(req.Images) > 0 {
		imageURLs, err := utils.UploadMultipleImagesWithValidation(req.Images)
		if err != nil {
			utils.CleanupImageOnError(req.ImageURL) // customize rollback upload cloudinary
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		req.ImageURLs = imageURLs
	}

	if err := h.service.CreateClass(req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		utils.CleanupImagesOnError(req.ImageURLs)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Class created successfully"})
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateClassRequest
	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	req.IsActive, _ = utils.ParseBoolFormField(c, "isActive")

	if req.Image != nil && req.Image.Filename != "" {
		imageURL, err := utils.UploadImageWithValidation(req.Image)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		req.ImageURL = imageURL
	}

	if err := h.service.UpdateClass(id, req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteClass(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}

func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id := c.Param("id")

	classResponse, err := h.service.GetClassByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Class not found"})
		return
	}
	c.JSON(http.StatusOK, classResponse)
}

func (h *ClassHandler) GetAllClasses(c *gin.Context) {
	var params dto.ClassQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	classes, pagination, err := h.service.GetAllClasses(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"classes":    classes,
		"pagination": pagination,
	})
}

func (h *ClassHandler) UploadClassGallery(c *gin.Context) {
	classID := c.Param("id")
	var req dto.UpdateGalleryRequest

	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	parsedID, err := uuid.Parse(classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid class ID"})
		return
	}

	imageURLs, err := utils.UploadMultipleImagesWithValidation(req.Images)
	if err != nil {
		utils.CleanupImagesOnError(imageURLs)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	req.ImageURLs = imageURLs

	if err := h.service.UpdateClassGallery(parsedID, req.KeepImages, req.ImageURLs); err != nil {
		utils.CleanupImagesOnError(req.ImageURLs)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gallery updated successfully"})
}
