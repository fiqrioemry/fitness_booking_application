package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	service services.ReviewService
}

func NewReviewHandler(service services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service}
}

func (h *ReviewHandler) CreateReviewFromBookingID(c *gin.Context) {
	bookingID := c.Param("id")
	var req dto.CreateReviewRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	userID := utils.MustGetUserID(c)

	if err := h.service.CreateReview(userID, bookingID, req); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Review created successfully"})
}

func (h *ReviewHandler) GetReviewsByClass(c *gin.Context) {
	classID := c.Param("classId")

	reviews, err := h.service.GetReviewsByClassID(classID)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, reviews)
}
