package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository         UserRepository
	AuthRepository         AuthRepository
	TypeRepository         TypeRepository
	ClassRepository        ClassRepository
	LevelRepository        LevelRepository
	ReviewRepository       ReviewRepository
	PaymentRepository      PaymentRepository
	BookingRepository      BookingRepository
	VoucherRepository      VoucherRepository
	PackageRepository      PackageRepository
	CategoryRepository     CategoryRepository
	LocationRepository     LocationRepository
	DashboardRepository    DashboardRepository
	InstructorRepository   InstructorRepository
	ScheduleRepository     ClassScheduleRepository
	UserPackageRepository  UserPackageRepository
	SubcategoryRepository  SubcategoryRepository
	TemplateRepository     ScheduleTemplateRepository
	NotificationRepository NotificationRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:         NewUserRepository(db),
		AuthRepository:         NewAuthRepository(db),
		TypeRepository:         NewTypeRepository(db),
		ClassRepository:        NewClassRepository(db),
		LevelRepository:        NewLevelRepository(db),
		ReviewRepository:       NewReviewRepository(db),
		PaymentRepository:      NewPaymentRepository(db),
		BookingRepository:      NewBookingRepository(db),
		VoucherRepository:      NewVoucherRepository(db),
		PackageRepository:      NewPackageRepository(db),
		CategoryRepository:     NewCategoryRepository(db),
		LocationRepository:     NewLocationRepository(db),
		DashboardRepository:    NewDashboardRepository(db),
		InstructorRepository:   NewInstructorRepository(db),
		ScheduleRepository:     NewClassScheduleRepository(db),
		UserPackageRepository:  NewUserPackageRepository(db),
		SubcategoryRepository:  NewSubcategoryRepository(db),
		TemplateRepository:     NewScheduleTemplateRepository(db),
		NotificationRepository: NewNotificationRepository(db),
	}
}
