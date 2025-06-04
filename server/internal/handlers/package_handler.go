package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type PackageHandler struct {
	service services.PackageService
}

func NewPackageHandler(service services.PackageService) *PackageHandler {
	return &PackageHandler{service}
}

func (h *PackageHandler) CreatePackage(c *gin.Context) {
	var req dto.CreatePackageRequest
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

	if err := h.service.CreatePackage(req); err != nil {
		utils.CleanupImageOnError(imageURL)
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Package created successfully"})
}

func (h *PackageHandler) UpdatePackage(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdatePackageRequest
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

	if err := h.service.UpdatePackage(id, req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package updated successfully"})
}

func (h *PackageHandler) GetAllPackages(c *gin.Context) {
	var params dto.PackageQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	packages, pagination, err := h.service.GetAllPackages(params)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       packages,
		"pagination": pagination,
	})
}

func (h *PackageHandler) GetPackageByID(c *gin.Context) {
	id := c.Param("id")

	classPackage, err := h.service.GetPackageByID(id)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, classPackage)
}

func (h *PackageHandler) DeletePackage(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeletePackage(id); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package deleted successfully"})
}
