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
	userService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(userService)

	// Route Binding ==========================
	routes.AuthRoutes(r, authHandler)

	// Start Server ===========================
	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
