package routes

import (
	"server/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, h *bootstrap.Handlers) {
	api := r.Group("/api/v1")

	// ========= Authentication & User Management ========
	AuthRoutes(api, h.AuthHandler)
	UserRoutes(api, h.UserHandler)

	// ========= Class & Schedule Management =============
	TypeRoutes(api, h.TypeHandler)
	LevelRoutes(api, h.LevelHandler)
	ClassRoutes(api, h.ClassHandler)
	LocationRoutes(api, h.LocationHandler)
	ScheduleRoutes(api, h.ScheduleHandler)
	TemplateRoutes(api, h.TemplateHandler)
	CategoryRoutes(api, h.CategoryHandler)
	InstructorRoutes(api, h.InstructorHandler)
	SubcategoryRoutes(api, h.SubcategoryHandler)

	// ======== Booking Management =======================
	ReviewRoutes(api, h.ReviewHandler)
	BookingRoutes(api, h.BookingHandler)
	PackageRoutes(api, h.PackageHandler)
	UserPackageRoutes(api, h.UserPackageHandler)

	// ======== Payment & Voucher ========================
	VoucherRoutes(api, h.VoucherHandler)
	PaymentRoutes(api, h.PaymentHandler)

	// ======== Notification & Dashboard =================
	NotificationRoutes(api, h.NotificationHandler)
	DashboardRoutes(api, h.DashboardHandler)
}
