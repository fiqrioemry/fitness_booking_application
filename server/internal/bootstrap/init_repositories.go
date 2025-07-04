package bootstrap

import (
	"server/internal/repositories"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository         repositories.UserRepository
	AuthRepository         repositories.AuthRepository
	TypeRepository         repositories.TypeRepository
	ClassRepository        repositories.ClassRepository
	LevelRepository        repositories.LevelRepository
	ReviewRepository       repositories.ReviewRepository
	PaymentRepository      repositories.PaymentRepository
	BookingRepository      repositories.BookingRepository
	VoucherRepository      repositories.VoucherRepository
	PackageRepository      repositories.PackageRepository
	CategoryRepository     repositories.CategoryRepository
	LocationRepository     repositories.LocationRepository
	DashboardRepository    repositories.DashboardRepository
	InstructorRepository   repositories.InstructorRepository
	ScheduleRepository     repositories.ClassScheduleRepository
	UserPackageRepository  repositories.UserPackageRepository
	SubcategoryRepository  repositories.SubcategoryRepository
	TemplateRepository     repositories.ScheduleTemplateRepository
	NotificationRepository repositories.NotificationRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:         repositories.NewUserRepository(db),
		AuthRepository:         repositories.NewAuthRepository(db),
		TypeRepository:         repositories.NewTypeRepository(db),
		ClassRepository:        repositories.NewClassRepository(db),
		LevelRepository:        repositories.NewLevelRepository(db),
		ReviewRepository:       repositories.NewReviewRepository(db),
		PaymentRepository:      repositories.NewPaymentRepository(db),
		BookingRepository:      repositories.NewBookingRepository(db),
		VoucherRepository:      repositories.NewVoucherRepository(db),
		PackageRepository:      repositories.NewPackageRepository(db),
		CategoryRepository:     repositories.NewCategoryRepository(db),
		LocationRepository:     repositories.NewLocationRepository(db),
		DashboardRepository:    repositories.NewDashboardRepository(db),
		InstructorRepository:   repositories.NewInstructorRepository(db),
		ScheduleRepository:     repositories.NewClassScheduleRepository(db),
		UserPackageRepository:  repositories.NewUserPackageRepository(db),
		SubcategoryRepository:  repositories.NewSubcategoryRepository(db),
		TemplateRepository:     repositories.NewScheduleTemplateRepository(db),
		NotificationRepository: repositories.NewNotificationRepository(db),
	}
}
