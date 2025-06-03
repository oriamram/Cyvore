package main

import (
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/internal/api/handler"
	"backend/internal/config"
	"backend/internal/middleware"
)

func main() {
	// Get configuration
	cfg := config.Get()

	// Setup router
	r := gin.Default()

	// Add CORS middleware with specific configuration
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000" || strings.HasPrefix(origin, "http://localhost:")
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public routes
	r.POST("/auth/register", handler.Register)
	r.POST("/auth/signin", handler.SignIn)
	r.POST("/auth/signout", handler.SignOut)
	r.POST("/auth/refresh", handler.RefreshToken)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/scan", handler.ScanHandler)
		protected.POST("/scan/stop", handler.StopScanHandler)
		protected.GET("/scan/status", handler.StatusHandler)
		protected.POST("/clean", handler.CleanHandler)
		// protected.GET("/amass/data", handler.AmassDataHandler)
	}

	// Start server
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}