package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, handler *handlers.PaymentHandler) {
	r.POST("/api/payments/stripe/notifications", handler.HandlePaymentNotifications)

	payments := r.Group("/api/payments")
	payments.Use(middleware.AuthRequired())
	payments.POST("", handler.CreatePayment)
	payments.GET("/me", handler.GetMyTransactions)
	payments.GET("/me/:id", handler.GetPaymentDetail)

	admin := payments.Use(middleware.RoleOnly("admin"))
	admin.GET("", handler.GetAllUserPayments)
	admin.GET("/:id", handler.GetPaymentDetail)

}
