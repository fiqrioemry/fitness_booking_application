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
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", username, password, host, port)
	dbRoot, err := gorm.Open(mysql.Open(dsnRoot), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to MySQL server: " + err.Error())
	}

	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", database)
	if err := dbRoot.Exec(sql).Error; err != nil {
		panic("Failed to create database: " + err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)

	for range 10 {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
