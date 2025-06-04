package seeders

import (
	"log"
	"server/internal/models"

	"gorm.io/gorm"
)

func ResetDatabase(db *gorm.DB) {
	log.Println("⚠️ Dropping all tables...")

	err := db.Migrator().DropTable(
		&models.Token{},
		&models.User{},
		&models.Package{},
		&models.PackageClass{},
		&models.UserPackage{},
		&models.Category{},
		&models.Subcategory{},
		&models.Type{},
		&models.Level{},
		&models.Class{},
		&models.ClassGallery{},
		&models.ClassSchedule{},
		&models.ScheduleTemplate{},
		&models.Booking{},
		&models.Payment{},
		&models.Notification{},
		&models.NotificationType{},
		&models.NotificationSetting{},
		&models.Voucher{},
		&models.UsedVoucher{},
		&models.Review{},
		&models.Attendance{},
		&models.Instructor{},
		&models.Location{},
	)
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	log.Println("All tables dropped successfully.")

	log.Println("migrating tables...")

	err = db.AutoMigrate(
		&models.Token{},
		&models.User{},
		&models.Package{},
		&models.PackageClass{},
		&models.UserPackage{},
		&models.Category{},
		&models.Subcategory{},
		&models.Type{},
		&models.Level{},
		&models.Class{},
		&models.ClassGallery{},
		&models.ClassSchedule{},
		&models.ScheduleTemplate{},
		&models.Booking{},
		&models.Payment{},
		&models.Notification{},
		&models.NotificationType{},
		&models.NotificationSetting{},
		&models.Voucher{},
		&models.UsedVoucher{},
		&models.Review{},
		&models.Attendance{},
		&models.Instructor{},
		&models.Location{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	log.Println("migration completed successfully.")

	log.Println("seeding dummy data...")

	SeedUsers(db)
	SeedCategories(db)
	SeedSubcategories(db)
	SeedTypes(db)
	SeedLevels(db)
	SeedLocations(db)
	SeedClasses(db)
	SeedClassGalleries(db)
	SeedInstructors(db)

	log.Println("seeding completed successfully.")
}
