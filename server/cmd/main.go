package main

import (
	"log"
	"os"
	"server/internal/config"
	"server/internal/handlers"
	"server/internal/repositories"
	"server/internal/routes"
	"server/internal/services"
	"server/pkg/middleware"
	"server/pkg/utils"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	// config =================================
	config.InitConfiguration()
	utils.InitLogger()

	// seeders
	db := config.DB
	// seeders.ResetDatabase(db)

	r := gin.Default()
	err := r.SetTrustedProxies(config.GetTrustedProxies())
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// middleware ======================
	r.Use(
		ginzap.Ginzap(utils.GetLogger(), time.RFC3339, true),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(20, 60*time.Second),
		middleware.LimitFileSize(12<<20),
		middleware.APIKeyGateway([]string{"/api/auth/google", "/api/auth/google/callback"}),
	)
	// Inisialisasi Dependency for auth module
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo, userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	classRepo := repositories.NewClassRepository(db)
	classService := services.NewClassService(classRepo)
	classHandler := handlers.NewClassHandler(classService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	typeRepo := repositories.NewTypeRepository(db)
	typeService := services.NewTypeService(typeRepo)
	typeHandler := handlers.NewTypeHandler(typeService)

	levelRepo := repositories.NewLevelRepository(db)
	levelService := services.NewLevelService(levelRepo)
	levelHandler := handlers.NewLevelHandler(levelService)

	locationRepo := repositories.NewLocationRepository(db)
	locationService := services.NewLocationService(locationRepo)
	locationHandler := handlers.NewLocationHandler(locationService)

	packageRepo := repositories.NewPackageRepository(db)
	packageService := services.NewPackageService(packageRepo)
	packageHandler := handlers.NewPackageHandler(packageService)

	subcategoryRepo := repositories.NewSubcategoryRepository(db)
	subcategoryService := services.NewSubcategoryService(subcategoryRepo)
	subcategoryHandler := handlers.NewSubcategoryHandler(subcategoryService)

	voucherRepo := repositories.NewVoucherRepository(db)
	voucherService := services.NewVoucherService(voucherRepo)
	voucherHandler := handlers.NewVoucherHandler(voucherService)

	instructorRepo := repositories.NewInstructorRepository(db)
	instructorService := services.NewInstructorService(instructorRepo, userRepo)
	instructorHandler := handlers.NewInstructorHandler(instructorService)

	notificationRepo := repositories.NewNotificationRepository(db)
	notificationService := services.NewNotificationService(notificationRepo)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	userPackageRepo := repositories.NewUserPackageRepository(db)
	userPackageService := services.NewUserPackageService(userPackageRepo)
	userPackageHandler := handlers.NewUserPackageHandler(userPackageService)

	paymentRepo := repositories.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(paymentRepo, packageRepo, userRepo, voucherService, notificationService, userPackageRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	bookingRepo := repositories.NewBookingRepository(db)
	bookingService := services.NewBookingService(bookingRepo)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	reviewRepo := repositories.NewReviewRepository(db)
	reviewService := services.NewReviewService(reviewRepo, bookingRepo, instructorRepo)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	// Route Binding ==========================
	routes.AuthRoutes(r, authHandler)
	routes.UserRoutes(r, userHandler)
	routes.TypeRoutes(r, typeHandler)
	routes.ClassRoutes(r, classHandler)
	routes.LevelRoutes(r, levelHandler)
	routes.PackageRoutes(r, packageHandler)
	routes.CategoryRoutes(r, categoryHandler)
	routes.VoucherRoutes(r, voucherHandler)
	routes.PaymentRoutes(r, paymentHandler)
	routes.LocationRoutes(r, locationHandler)
	routes.InstructorRoutes(r, instructorHandler)
	routes.SubcategoryRoutes(r, subcategoryHandler)
	routes.NotificationRoutes(r, notificationHandler)
	routes.UserPackageRoutes(r, userPackageHandler)
	routes.ReviewRoutes(r, reviewHandler)
	routes.BookingRoutes(r, bookingHandler)

	// Start Server ===========================
	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
