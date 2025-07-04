package bootstrap

import "server/internal/handlers"

type Handlers struct {
	AuthHandler         *handlers.AuthHandler
	UserHandler         *handlers.UserHandler
	TypeHandler         *handlers.TypeHandler
	LevelHandler        *handlers.LevelHandler
	ClassHandler        *handlers.ClassHandler
	ReviewHandler       *handlers.ReviewHandler
	PackageHandler      *handlers.PackageHandler
	VoucherHandler      *handlers.VoucherHandler
	PaymentHandler      *handlers.PaymentHandler
	BookingHandler      *handlers.BookingHandler
	LocationHandler     *handlers.LocationHandler
	CategoryHandler     *handlers.CategoryHandler
	DashboardHandler    *handlers.DashboardHandler
	ScheduleHandler     *handlers.ClassScheduleHandler
	InstructorHandler   *handlers.InstructorHandler
	TemplateHandler     *handlers.ScheduleTemplateHandler
	UserPackageHandler  *handlers.UserPackageHandler
	SubcategoryHandler  *handlers.SubcategoryHandler
	NotificationHandler *handlers.NotificationHandler
}

func InitHandlers(s *Services) *Handlers {
	return &Handlers{
		AuthHandler:         handlers.NewAuthHandler(s.AuthService),
		UserHandler:         handlers.NewUserHandler(s.UserService),
		TypeHandler:         handlers.NewTypeHandler(s.TypeService),
		LevelHandler:        handlers.NewLevelHandler(s.LevelService),
		ClassHandler:        handlers.NewClassHandler(s.ClassService),
		ReviewHandler:       handlers.NewReviewHandler(s.ReviewService),
		PackageHandler:      handlers.NewPackageHandler(s.PackageService),
		VoucherHandler:      handlers.NewVoucherHandler(s.VoucherService),
		PaymentHandler:      handlers.NewPaymentHandler(s.PaymentService),
		BookingHandler:      handlers.NewBookingHandler(s.BookingService),
		LocationHandler:     handlers.NewLocationHandler(s.LocationService),
		CategoryHandler:     handlers.NewCategoryHandler(s.CategoryService),
		DashboardHandler:    handlers.NewDashboardHandler(s.DashboardService),
		ScheduleHandler:     handlers.NewClassScheduleHandler(s.ScheduleService),
		InstructorHandler:   handlers.NewInstructorHandler(s.InstructorService),
		TemplateHandler:     handlers.NewScheduleTemplateHandler(s.TemplateService),
		UserPackageHandler:  handlers.NewUserPackageHandler(s.UserPackageService),
		SubcategoryHandler:  handlers.NewSubcategoryHandler(s.SubcategoryService),
		NotificationHandler: handlers.NewNotificationHandler(s.NotificationService),
	}
}
