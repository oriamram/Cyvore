package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"backend/internal/api/handler"
	"backend/internal/config"
)

func main() {
    // Load configuration
    cfg := config.Load()
    
    // Setup router
    r := gin.Default()

    // Routes
    r.GET("/scan", handler.ScanHandler)
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    // Start server
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}