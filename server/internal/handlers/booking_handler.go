package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	bookingService services.BookingService
}

func NewBookingHandler(bookingService services.BookingService) *BookingHandler {
	return &BookingHandler{bookingService}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.CreateBookingRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	err := h.bookingService.CreateBooking(userID, req.PackageID, req.ClassScheduleID)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking successful"})
}

func (h *BookingHandler) GetMyBookings(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var params dto.BookingQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	data, pagination, err := h.bookingService.GetBookingByUser(userID, params)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       data,
		"pagination": pagination,
	})
}

func (h *BookingHandler) GetBookingDetail(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	bookingID := c.Param("id")

	result, err := h.bookingService.GetBookingDetail(userID, bookingID)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *BookingHandler) CheckinBookedClass(c *gin.Context) {
	bookingID := c.Param("id")
	userID := utils.MustGetUserID(c)

	err := h.bookingService.CheckedInClassSchedule(userID, bookingID)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checkin Successfully"})
}

func (h *BookingHandler) CheckoutBookedClass(c *gin.Context) {
	bookingID := c.Param("id")
	userID := utils.MustGetUserID(c)

	var req dto.ValidateCheckoutRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.bookingService.CheckoutClassSchedule(userID, bookingID, req); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance verified successfully"})
}
