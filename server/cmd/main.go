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

	db := config.DB
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
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	classRepo := repositories.NewClassRepository(db)
	classService := services.NewClassService(classRepo)
	classHandler := handlers.NewClassHandler(classService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	levelRepo := repositories.NewLevelRepository(db)
	levelService := services.NewLevelService(levelRepo)
	levelHandler := handlers.NewLevelHandler(levelService)

	locationRepo := repositories.NewLocationRepository(db)
	locationService := services.NewLocationService(locationRepo)
	locationHandler := handlers.NewLocationHandler(locationService)

	subcategoryRepo := repositories.NewSubcategoryRepository(db)
	subcategoryService := services.NewSubcategoryService(subcategoryRepo)
	subcategoryHandler := handlers.NewSubcategoryHandler(subcategoryService)

	// Route Binding ==========================
	routes.AuthRoutes(r, authHandler)
	routes.UserRoutes(r, userHandler)
	routes.ClassRoutes(r, classHandler)
	routes.LevelRoutes(r, levelHandler)
	routes.CategoryRoutes(r, categoryHandler)
	routes.LocationRoutes(r, locationHandler)
	routes.SubcategoryRoutes(r, subcategoryHandler)

	// Start Server ===========================
	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
