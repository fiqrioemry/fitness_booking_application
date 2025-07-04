package bootstrap

import (
	"server/internal/services"

	"gorm.io/gorm"
)

type Services struct {
	UserService         services.UserService
	AuthService         services.AuthService
	TypeService         services.TypeService
	ClassService        services.ClassService
	LevelService        services.LevelService
	ReviewService       services.ReviewService
	PaymentService      services.PaymentService
	BookingService      services.BookingService
	VoucherService      services.VoucherService
	PackageService      services.PackageService
	CategoryService     services.CategoryService
	LocationService     services.LocationService
	DashboardService    services.DashboardService
	InstructorService   services.InstructorService
	ScheduleService     services.ClassScheduleService
	UserPackageService  services.UserPackageService
	SubcategoryService  services.SubcategoryService
	TemplateService     services.ScheduleTemplateService
	NotificationService services.NotificationService
}

func InitServices(r *Repositories, db *gorm.DB) *Services {
	notificationService := services.NewNotificationService(r.NotificationRepository)
	voucherService := services.NewVoucherService(r.VoucherRepository)
	templateService := services.NewScheduleTemplateService(
		r.TemplateRepository, r.ClassRepository, r.InstructorRepository, r.ScheduleRepository,
	)

	return &Services{
		UserService:         services.NewUserService(r.UserRepository),
		AuthService:         services.NewAuthService(r.AuthRepository, r.UserRepository, r.NotificationRepository),
		TypeService:         services.NewTypeService(r.TypeRepository),
		ClassService:        services.NewClassService(r.ClassRepository),
		LevelService:        services.NewLevelService(r.LevelRepository),
		ReviewService:       services.NewReviewService(r.ReviewRepository, r.BookingRepository, r.InstructorRepository),
		PaymentService:      services.NewPaymentService(r.PaymentRepository, r.PackageRepository, r.UserRepository, voucherService, notificationService, r.UserPackageRepository),
		BookingService:      services.NewBookingService(db, r.BookingRepository, r.PackageRepository, notificationService, r.UserPackageRepository, r.ScheduleRepository),
		VoucherService:      voucherService,
		PackageService:      services.NewPackageService(r.PackageRepository),
		CategoryService:     services.NewCategoryService(r.CategoryRepository),
		LocationService:     services.NewLocationService(r.LocationRepository),
		DashboardService:    services.NewDashboardService(r.DashboardRepository),
		InstructorService:   services.NewInstructorService(r.InstructorRepository, r.UserRepository),
		ScheduleService:     services.NewClassScheduleService(r.ScheduleRepository, templateService, r.ClassRepository, r.InstructorRepository, r.BookingRepository, r.PackageRepository),
		UserPackageService:  services.NewUserPackageService(r.UserPackageRepository),
		SubcategoryService:  services.NewSubcategoryService(r.SubcategoryRepository),
		TemplateService:     templateService,
		NotificationService: notificationService,
	}
}
