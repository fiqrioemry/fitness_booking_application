package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(r *gin.Engine, h *handlers.ReviewHandler) {
	review := r.Group("/api/reviews")
	review.Use(middleware.AuthRequired())

	review.POST("/:id", h.CreateReviewFromBookingID)
	r.GET("/api/reviews/:classId", h.GetReviewsByClass)
}
