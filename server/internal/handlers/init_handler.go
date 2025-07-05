package handlers

import (
	"server/internal/services"
)

type Handlers struct {
	AuthHandler         *AuthHandler
	UserHandler         *UserHandler
	TypeHandler         *TypeHandler
	LevelHandler        *LevelHandler
	ClassHandler        *ClassHandler
	ReviewHandler       *ReviewHandler
	PackageHandler      *PackageHandler
	VoucherHandler      *VoucherHandler
	PaymentHandler      *PaymentHandler
	BookingHandler      *BookingHandler
	LocationHandler     *LocationHandler
	CategoryHandler     *CategoryHandler
	DashboardHandler    *DashboardHandler
	ScheduleHandler     *ClassScheduleHandler
	InstructorHandler   *InstructorHandler
	TemplateHandler     *ScheduleTemplateHandler
	UserPackageHandler  *UserPackageHandler
	SubcategoryHandler  *SubcategoryHandler
	NotificationHandler *NotificationHandler
}

func InitHandlers(s *services.Services) *Handlers {
	return &Handlers{
		AuthHandler:         NewAuthHandler(s.AuthService),
		UserHandler:         NewUserHandler(s.UserService),
		TypeHandler:         NewTypeHandler(s.TypeService),
		LevelHandler:        NewLevelHandler(s.LevelService),
		ClassHandler:        NewClassHandler(s.ClassService),
		ReviewHandler:       NewReviewHandler(s.ReviewService),
		PackageHandler:      NewPackageHandler(s.PackageService),
		VoucherHandler:      NewVoucherHandler(s.VoucherService),
		PaymentHandler:      NewPaymentHandler(s.PaymentService),
		BookingHandler:      NewBookingHandler(s.BookingService),
		LocationHandler:     NewLocationHandler(s.LocationService),
		CategoryHandler:     NewCategoryHandler(s.CategoryService),
		DashboardHandler:    NewDashboardHandler(s.DashboardService),
		ScheduleHandler:     NewClassScheduleHandler(s.ScheduleService),
		InstructorHandler:   NewInstructorHandler(s.InstructorService),
		TemplateHandler:     NewScheduleTemplateHandler(s.TemplateService),
		UserPackageHandler:  NewUserPackageHandler(s.UserPackageService),
		SubcategoryHandler:  NewSubcategoryHandler(s.SubcategoryService),
		NotificationHandler: NewNotificationHandler(s.NotificationService),
	}
}
