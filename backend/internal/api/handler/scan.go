package handler

import (
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func ScanHandler(c *gin.Context) {
	domain := c.Query("domain")
	if domain == "" {
		c.JSON(400, gin.H{"error": "domain parameter is required"})
		return
	}

	// Get the scan service instance
	scanService, err := service.GetScanService()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to initialize scan service: " + err.Error()})
		return
	}

	// Start the scan
	err = scanService.GetDomainInfo(domain)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to start scan: " + err.Error()})
		return
	}

	// Return success immediately
	c.JSON(200, gin.H{
		"message": "Scan started successfully",
	})
}

// StopScanHandler handles requests to stop the running scan
func StopScanHandler(c *gin.Context) {
	// Get the scan service instance
	scanService, err := service.GetScanService()
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error": "Failed to initialize scan service: " + err.Error(),
		})
		return
	}

	// Stop the scan
	err = scanService.StopScan()
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error": "Failed to stop scan: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Scan stopped successfully",
	})
}

// StatusHandler returns the current scan status
func StatusHandler(c *gin.Context) {
	// Get the scan service instance
	scanService, err := service.GetScanService()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to initialize scan service: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"scanning": scanService.IsScanning(),
	})
}