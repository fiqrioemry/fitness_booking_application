// internal/routes/review_route.go
package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(r *gin.Engine, h *handlers.ReviewHandler) {
	// public-endpoints
	r.GET("/api/reviews/:classId", h.GetReviewsByClass)

	review := r.Group("/api/reviews")
	// customer-endpoints
	review.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	review.POST("/:id", h.CreateReviewFromBookingID)
}
