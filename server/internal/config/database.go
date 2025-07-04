package config

import (
	"fmt"
	"os"
	"time"

	"server/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dbRootURL := os.Getenv("DB_ROOT_URL")
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	// create root connection
	dbRoot, err := gorm.Open(mysql.Open(dbRootURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to MySQL server: " + err.Error())
	}

	// Create database if not exists
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	if err := dbRoot.Exec(sql).Error; err != nil {
		panic("Failed to create database: " + err.Error())
	}

	// connect to the specific database
	for range 10 {
		DB, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("Waiting for database to be ready...")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// migration
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Package{},
		&models.PackageClass{},
		&models.UserPackage{},
		&models.Category{},
		&models.Subcategory{},
		&models.Type{},
		&models.Level{},
		&models.Class{},
		&models.Location{},
		&models.Instructor{},
		&models.ClassGallery{},
		&models.ClassSchedule{},
		&models.ScheduleTemplate{},
		&models.Payment{},
		&models.Review{},
		&models.Booking{},
		&models.Attendance{},
		&models.Voucher{},
		&models.UsedVoucher{},
		&models.Notification{},
		&models.NotificationType{},
		&models.NotificationSetting{},
	); err != nil {
		panic("Migration failed: " + err.Error())
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get database connection: " + err.Error())
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Database connection established successfully.")
}
