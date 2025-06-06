package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine, h *handlers.BookingHandler) {
	customer := r.Group("/api/bookings")
	// customer-endpoint
	customer.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	customer.POST("", h.CreateBooking)
	customer.GET("", h.GetMyBookings)
	customer.GET("/:id", h.GetBookingDetail)
	customer.POST("/:id/check-in", h.CheckinBookedClass)
	customer.POST("/:id/check-out", h.CheckoutBookedClass)
}
