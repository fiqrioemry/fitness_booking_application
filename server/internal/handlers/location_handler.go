package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	locationService services.LocationService
}

func NewLocationHandler(locationService services.LocationService) *LocationHandler {
	return &LocationHandler{locationService}
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var req dto.CreateLocationRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.locationService.CreateLocation(req); err != nil {
		utils.HandleServiceError(c, err, "Failed to create location")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Location created successfully"})
}

func (h *LocationHandler) UpdateLocation(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateLocationRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.locationService.UpdateLocation(id, req); err != nil {
		utils.HandleServiceError(c, err, "Failed to update location")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location updated successfully"})
}

func (h *LocationHandler) DeleteLocation(c *gin.Context) {
	id := c.Param("id")

	if err := h.locationService.DeleteLocation(id); err != nil {
		utils.HandleServiceError(c, err, "Failed to delete location")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}

func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	locations, err := h.locationService.GetAllLocations()
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch locations")
		return
	}

	if locations == nil {
		locations = []dto.LocationResponse{}
	}

	c.JSON(http.StatusOK, locations)
}

func (h *LocationHandler) GetLocationByID(c *gin.Context) {
	id := c.Param("id")

	location, err := h.locationService.GetLocationByID(id)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch location")
		return
	}

	c.JSON(http.StatusOK, location)
}
