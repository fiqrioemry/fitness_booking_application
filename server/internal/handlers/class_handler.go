package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
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

	req.IsActive = utils.ParseBoolFormField(c, "isActive")

	imageURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	req.ImageURL = imageURL

	if len(req.Images) > 0 {
		imageURLs, err := utils.UploadMultipleImagesWithValidation(req.Images)
		if err != nil {
			utils.CleanupImageOnError(req.ImageURL)
			utils.HandleServiceError(c, err, err.Error())
			return
		}
		req.ImageURLs = imageURLs
	}

	if err := h.service.CreateClass(req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		utils.CleanupImagesOnError(req.ImageURLs)
		utils.HandleServiceError(c, err, "Failed to create class")
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

	req.IsActive = utils.ParseBoolFormField(c, "isActive")

	if req.Image != nil && req.Image.Filename != "" {
		imageURL, err := utils.UploadImageWithValidation(req.Image)
		if err != nil {
			utils.HandleServiceError(c, err, err.Error())
			return
		}
		req.ImageURL = imageURL
	}

	if err := h.service.UpdateClass(id, req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		utils.HandleServiceError(c, err, "Failed to update class")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteClass(id); err != nil {
		utils.HandleServiceError(c, err, "Failed to delete class")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}

func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id := c.Param("id")

	classResponse, err := h.service.GetClassByID(id)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to get class")
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
		utils.HandleServiceError(c, err, "Failed to get all classes")
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

	form := utils.GetMultipartForm(c)
	if c.IsAborted() {
		return
	}

	oldImages, newUpload := utils.ParseMixedImages(form)
	parsedID := utils.MustParseUUID(c, classID, "classID")

	newImageURLs, err := utils.UploadMultipleImagesWithValidation(newUpload)
	if err != nil {
		utils.CleanupImagesOnError(newImageURLs)
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	req.ImageURLs = append(req.ImageURLs, newImageURLs...)

	if err := h.service.UpdateClassGallery(parsedID, oldImages, req.ImageURLs); err != nil {
		utils.CleanupImagesOnError(req.ImageURLs)
		utils.HandleServiceError(c, err, "Failed to update class gallery")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gallery updated successfully"})
}
