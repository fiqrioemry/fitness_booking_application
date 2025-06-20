package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.RouterGroup, h *handlers.PaymentHandler) {
	// webhook-endpoints
	r.POST("/payments/stripe/notifications", h.HandlePaymentNotifications)

	// customer-endpoints
	customer := r.Group("/payments")
	customer.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	customer.POST("", h.CreatePayment)
	customer.GET("/me", h.GetMyTransactions)
	customer.GET("/me/:id", h.GetPaymentDetail)

	// admin-endpoints
	admin := r.Group("/admin/payments")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.GET("", h.GetAllUserPayments)
	admin.GET("/:id", h.GetPaymentDetail)
}
