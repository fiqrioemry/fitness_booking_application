package main

import (
	"log"
	"net/http"
	"os"

	"server/internal/config"
	"server/internal/cron"
	"server/internal/handlers"
	"server/internal/repositories"
	"server/internal/routes"
	"server/internal/seeders"
	"server/internal/services"
	"server/pkg/middleware"
	"server/pkg/utils"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	// ========== Configuration =================
	config.InitConfiguration()
	utils.InitLogger()
	db := config.DB

	seeders.ResetDatabase(db)

	// ========== initialisasi layer ============
	repo := repositories.InitRepositories(db)
	s := services.InitServices(repo, db)
	h := handlers.InitHandlers(s)

	// ========== inisialisasi cron job =========
	cronManager := cron.NewCronManager(s.PaymentService, s.TemplateService, s.NotificationService, s.BookingService)
	cronManager.RegisterJobs()
	cronManager.Start()

	// ========== Inisialisasi gin engine =======
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		log.Println("Root endpoint accessed just now", gin.H{
			"ip":        c.ClientIP(),
			"userAgent": c.GetHeader("User-Agent"),
		})

		c.JSON(http.StatusOK, gin.H{
			"status":    "success",
			"message":   "Welcome to sweatup API",
			"version":   "1.4.0",
			"ip":        c.ClientIP(),
			"userAgent": c.GetHeader("User-Agent"),
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "healthy",
			"timestamp": utils.NowISO(),
			"uptime":    utils.GetUptime(),
		})
	})

	if err := r.SetTrustedProxies(config.GetTrustedProxies()); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// ========== inisialisasi Middleware ========
	r.Use(
		ginzap.Ginzap(utils.GetLogger(), time.RFC3339, true),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(100, 60*time.Second),
		middleware.LimitFileSize(12<<20),
		middleware.APIKeyGateway([]string{
			"/api/v1/auth/google",
			"/api/v1/auth/google/callback",
		}),
	)

	// ========== inisialisasi routes ===========
	routes.InitRoutes(r, h)

	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
