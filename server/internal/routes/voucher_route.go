package routes

import (
	"server/internal/handlers"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func VoucherRoutes(r *gin.RouterGroup, h *handlers.VoucherHandler) {
	// public-endpoints
	voucher := r.Group("/vouchers")
	voucher.POST("/apply", h.ApplyVoucher)
	voucher.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	voucher.GET("", h.GetAllVouchers)

	// admin-endpoints
	admin := r.Group("/admin/vouchers")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateVoucher)
	admin.PUT("/:id", h.UpdateVoucher)
	admin.DELETE("/:id", h.DeleteVoucher)
}
