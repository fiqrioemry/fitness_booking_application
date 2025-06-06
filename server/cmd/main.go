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
	// ========== Configuration ==========
	config.InitConfiguration()
	utils.InitLogger()
	db := config.DB

	// ========== Gin Engine ==========
	r := gin.Default()
	if err := r.SetTrustedProxies(config.GetTrustedProxies()); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// ========== Middleware ==========
	r.Use(
		ginzap.Ginzap(utils.GetLogger(), time.RFC3339, true),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(100, 60*time.Second),
		middleware.LimitFileSize(12<<20),
		middleware.APIKeyGateway([]string{"/api/auth/google", "/api/auth/google/callback"}),
	)

	// ========== Repositories ==========
	userRepo := repositories.NewUserRepository(db)
	authRepo := repositories.NewAuthRepository(db)
	typeRepo := repositories.NewTypeRepository(db)
	classRepo := repositories.NewClassRepository(db)
	levelRepo := repositories.NewLevelRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)
	paymentRepo := repositories.NewPaymentRepository(db)
	bookingRepo := repositories.NewBookingRepository(db)
	voucherRepo := repositories.NewVoucherRepository(db)
	packageRepo := repositories.NewPackageRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	locationRepo := repositories.NewLocationRepository(db)
	dashboardRepo := repositories.NewDashboardRepository(db)
	instructorRepo := repositories.NewInstructorRepository(db)
	scheduleRepo := repositories.NewClassScheduleRepository(db)
	userPackageRepo := repositories.NewUserPackageRepository(db)
	subcategoryRepo := repositories.NewSubcategoryRepository(db)
	templateRepo := repositories.NewScheduleTemplateRepository(db)
	notificationRepo := repositories.NewNotificationRepository(db)

	// ========== inisialisasi services ==========
	userService := services.NewUserService(userRepo)
	typeService := services.NewTypeService(typeRepo)
	levelService := services.NewLevelService(levelRepo)
	classService := services.NewClassService(classRepo)
	packageService := services.NewPackageService(packageRepo)
	voucherService := services.NewVoucherService(voucherRepo)
	authService := services.NewAuthService(authRepo, userRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	locationService := services.NewLocationService(locationRepo)
	dashboardService := services.NewDashboardService(dashboardRepo)
	userPackageService := services.NewUserPackageService(userPackageRepo)
	subcategoryService := services.NewSubcategoryService(subcategoryRepo)
	notificationService := services.NewNotificationService(notificationRepo)
	instructorService := services.NewInstructorService(instructorRepo, userRepo)
	reviewService := services.NewReviewService(reviewRepo, bookingRepo, instructorRepo)
	templateService := services.NewScheduleTemplateService(templateRepo, classRepo, instructorRepo, scheduleRepo)
	bookingService := services.NewBookingService(db, bookingRepo, packageRepo, notificationService, userPackageRepo, scheduleRepo)
	paymentService := services.NewPaymentService(paymentRepo, packageRepo, userRepo, voucherService, notificationService, userPackageRepo)
	scheduleService := services.NewClassScheduleService(scheduleRepo, templateService, classRepo, instructorRepo, bookingRepo, packageRepo)

	// ========== inisialisasi handlers ==========
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	typeHandler := handlers.NewTypeHandler(typeService)
	levelHandler := handlers.NewLevelHandler(levelService)
	classHandler := handlers.NewClassHandler(classService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	packageHandler := handlers.NewPackageHandler(packageService)
	voucherHandler := handlers.NewVoucherHandler(voucherService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	locationHandler := handlers.NewLocationHandler(locationService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)
	scheduleHandler := handlers.NewClassScheduleHandler(scheduleService)
	instructorHandler := handlers.NewInstructorHandler(instructorService)
	templateHandler := handlers.NewScheduleTemplateHandler(templateService)
	userPackageHandler := handlers.NewUserPackageHandler(userPackageService)
	subcategoryHandler := handlers.NewSubcategoryHandler(subcategoryService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	// ========== inisialisasi routes ==========
	routes.AuthRoutes(r, authHandler)
	routes.UserRoutes(r, userHandler)
	routes.TypeRoutes(r, typeHandler)
	routes.ClassRoutes(r, classHandler)
	routes.LevelRoutes(r, levelHandler)
	routes.ReviewRoutes(r, reviewHandler)
	routes.PackageRoutes(r, packageHandler)
	routes.VoucherRoutes(r, voucherHandler)
	routes.PaymentRoutes(r, paymentHandler)
	routes.BookingRoutes(r, bookingHandler)
	routes.LocationRoutes(r, locationHandler)
	routes.CategoryRoutes(r, categoryHandler)
	routes.TemplateRoutes(r, templateHandler)
	routes.ScheduleRoutes(r, scheduleHandler)
	routes.DashboardRoutes(r, dashboardHandler)
	routes.InstructorRoutes(r, instructorHandler)
	routes.UserPackageRoutes(r, userPackageHandler)
	routes.SubcategoryRoutes(r, subcategoryHandler)
	routes.NotificationRoutes(r, notificationHandler)

	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
