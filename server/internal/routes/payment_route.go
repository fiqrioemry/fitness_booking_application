package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, h *handlers.PaymentHandler) {
	r.POST("/api/payments/stripe/notifications", h.HandlePaymentNotifications)

	payments := r.Group("/api/payments")
	payments.Use(middleware.AuthRequired())
	payments.POST("", h.CreatePayment)
	payments.GET("/me", h.GetMyTransactions)
	payments.GET("/me/:id", h.GetPaymentDetail)

	admin := payments.Use(middleware.RoleOnly("admin"))
	admin.GET("", h.GetAllUserPayments)
	admin.GET("/:id", h.GetPaymentDetail)

}
