package handlers

import (
	"net/http"
	"os"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75/webhook"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.CreatePaymentRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	response, err := h.paymentService.CreatePayment(userID, req)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to create payment")
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *PaymentHandler) HandlePaymentNotifications(c *gin.Context) {
	const MaxBodyBytes = int64(65536)

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	body, err := c.GetRawData()
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to read request body")
		return
	}

	sigHeader := c.GetHeader("Stripe-Signature")

	event, err := webhook.ConstructEventWithOptions(body, sigHeader, os.Getenv("STRIPE_WEBHOOK_SECRET"), webhook.ConstructEventOptions{
		IgnoreAPIVersionMismatch: true,
	})
	if err != nil {
		utils.HandleServiceError(c, err, "Invalid Signatures")
		return
	}

	if err := h.paymentService.StripeWebhookNotification(event); err != nil {
		utils.HandleServiceError(c, err, "Failed to send notifications")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment received successfully"})
}

func (h *PaymentHandler) GetAllUserPayments(c *gin.Context) {
	var params dto.PaymentQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	payments, pagination, err := h.paymentService.GetAllUserPayments(params)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch payments")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       payments,
		"pagination": pagination,
	})

}

func (h *PaymentHandler) GetPaymentDetail(c *gin.Context) {
	paymentID := c.Param("id")
	payments, err := h.paymentService.GetPaymentDetail(paymentID)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to get payment")
		return
	}

	c.JSON(http.StatusOK, payments)

}

func (h *PaymentHandler) GetMyTransactions(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var params dto.PaymentQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}
	payments, pagination, err := h.paymentService.GetPaymentsByUserID(userID, params)
	if err != nil {
		utils.HandleServiceError(c, err, "Failed to fetch payments")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       payments,
		"pagination": pagination,
	})

}
