package services

import (
	"server/internal/repositories"

	"gorm.io/gorm"
)

type Services struct {
	UserService         UserService
	AuthService         AuthService
	TypeService         TypeService
	ClassService        ClassService
	LevelService        LevelService
	ReviewService       ReviewService
	PaymentService      PaymentService
	BookingService      BookingService
	VoucherService      VoucherService
	PackageService      PackageService
	CategoryService     CategoryService
	LocationService     LocationService
	DashboardService    DashboardService
	InstructorService   InstructorService
	ScheduleService     ClassScheduleService
	UserPackageService  UserPackageService
	SubcategoryService  SubcategoryService
	TemplateService     ScheduleTemplateService
	NotificationService NotificationService
}

func InitServices(r *repositories.Repositories, db *gorm.DB) *Services {
	notificationService := NewNotificationService(r.NotificationRepository)
	voucherService := NewVoucherService(r.VoucherRepository)
	templateService := NewScheduleTemplateService(
		r.TemplateRepository, r.ClassRepository, r.InstructorRepository, r.ScheduleRepository,
	)

	return &Services{
		UserService:         NewUserService(r.UserRepository),
		AuthService:         NewAuthService(r.AuthRepository, r.UserRepository, r.NotificationRepository),
		TypeService:         NewTypeService(r.TypeRepository),
		ClassService:        NewClassService(r.ClassRepository),
		LevelService:        NewLevelService(r.LevelRepository),
		ReviewService:       NewReviewService(r.ReviewRepository, r.BookingRepository, r.InstructorRepository),
		PaymentService:      NewPaymentService(r.PaymentRepository, r.PackageRepository, r.UserRepository, voucherService, notificationService, r.UserPackageRepository),
		BookingService:      NewBookingService(db, r.BookingRepository, r.PackageRepository, notificationService, r.UserPackageRepository, r.ScheduleRepository),
		VoucherService:      voucherService,
		PackageService:      NewPackageService(r.PackageRepository),
		CategoryService:     NewCategoryService(r.CategoryRepository),
		LocationService:     NewLocationService(r.LocationRepository),
		DashboardService:    NewDashboardService(r.DashboardRepository),
		InstructorService:   NewInstructorService(r.InstructorRepository, r.UserRepository),
		ScheduleService:     NewClassScheduleService(r.ScheduleRepository, templateService, r.ClassRepository, r.InstructorRepository, r.BookingRepository, r.PackageRepository),
		UserPackageService:  NewUserPackageService(r.UserPackageRepository),
		SubcategoryService:  NewSubcategoryService(r.SubcategoryRepository),
		TemplateService:     templateService,
		NotificationService: notificationService,
	}
}
