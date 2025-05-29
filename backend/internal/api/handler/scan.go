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

	// Initialize the scan service
	scanService, err := service.NewScanService()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to initialize scan service: " + err.Error()})
		return
	}

	// Run both scans
	enumOutput, err := scanService.GetDomainInfo(domain)
	if err != nil {
		c.JSON(500, gin.H{"error": "Enum scan failed: " + err.Error()})
		return
	}

	// Return both outputs as JSON
	c.JSON(200, gin.H{
		"enum":  enumOutput,
	})
}