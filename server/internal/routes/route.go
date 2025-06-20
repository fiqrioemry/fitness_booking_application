package routes

import (
	"server/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, h *bootstrap.Handlers) {
	api := r.Group("/api/v1")

	AuthRoutes(api, h.AuthHandler)
	UserRoutes(api, h.UserHandler)
	TypeRoutes(api, h.TypeHandler)
	LevelRoutes(api, h.LevelHandler)
	ClassRoutes(api, h.ClassHandler)
	ReviewRoutes(api, h.ReviewHandler)
	VoucherRoutes(api, h.VoucherHandler)
	PackageRoutes(api, h.PackageHandler)
	BookingRoutes(api, h.BookingHandler)
	PaymentRoutes(api, h.PaymentHandler)
	CategoryRoutes(api, h.CategoryHandler)
	LocationRoutes(api, h.LocationHandler)
	ScheduleRoutes(api, h.ScheduleHandler)
	TemplateRoutes(api, h.TemplateHandler)
	DashboardRoutes(api, h.DashboardHandler)
	InstructorRoutes(api, h.InstructorHandler)
	SubcategoryRoutes(api, h.SubcategoryHandler)
	UserPackageRoutes(api, h.UserPackageHandler)
	NotificationRoutes(api, h.NotificationHandler)

}
