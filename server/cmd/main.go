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

	db := config.DB
	r := gin.Default()

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

	// Route Binding ==========================
	routes.AuthRoutes(r, authHandler)
	routes.UserRoutes(r, userHandler)
	routes.ClassRoutes(r, classHandler)
	routes.CategoryRoutes(r, categoryHandler)

	// Start Server ===========================
	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
