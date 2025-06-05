package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine, h *handlers.BookingHandler) {
	booking := r.Group("/api/bookings")
	// customer-protected endpoints
	booking.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	booking.POST("", h.CreateBooking)
	booking.GET("", h.GetMyBookings)
	booking.GET("/:id", h.GetBookingDetail)
	booking.POST("/:id/check-in", h.CheckinBookedClass)
	booking.POST("/:id/check-out", h.CheckoutBookedClass)
}
