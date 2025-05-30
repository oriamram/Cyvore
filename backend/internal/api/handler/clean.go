package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"backend/internal/config"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

// CleanHandler handles requests to clean the database and logs
func CleanHandler(c *gin.Context) {
	cfg := config.Get()
	dataPath := cfg.DataPath

	// Clean database
	dbPath := filepath.Join(dataPath, "amass", "amass.sqlite")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to open database: %v", err)})
		return
	}
	defer db.Close()

	// Delete all data from tables
	_, err = db.Exec("DELETE FROM assets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to clean assets table: %v", err)})
		return
	}

	_, err = db.Exec("DELETE FROM relations")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to clean relations table: %v", err)})
		return
	}

	// Clean log files
	logPath := filepath.Join(dataPath, "amass", "log.txt")
	if err := os.Remove(logPath); err != nil && !os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to remove log file: %v", err)})
		return
	}

	resultsPath := filepath.Join(dataPath, "amass", "results.txt")
	if err := os.Remove(resultsPath); err != nil && !os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to remove results file: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Database and logs cleaned successfully",
	})
} 