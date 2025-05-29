package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"backend/internal/api/handler"
	"backend/internal/config"
)

func main() {
    // Get configuration
    cfg := config.Get()
    
    // Setup router
    r := gin.Default()

    // Routes
    r.GET("/scan", handler.ScanHandler)
    r.POST("/scan/stop", handler.StopScanHandler)
    r.GET("/scan/status", handler.StatusHandler)
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    // Start server
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}