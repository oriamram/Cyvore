package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
    Port     string
    DataPath string
    // Add other config fields as needed
    // DBHost string
    // JWTSecret string
}

// Global configuration instance
var cfg *Config

// Get returns the global configuration instance
func Get() *Config {
    if cfg == nil {
        cfg = Load()
    }
    return cfg
}

func Load() *Config {
    // Try to load .env file (don't fail if it doesn't exist)
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Println("No .env file found, using environment variables and defaults")
    }
    
    // Get the absolute path to the data directory
    dataPath := getEnv("DATA_PATH", "data")
    absDataPath, err := filepath.Abs(dataPath)
    if err != nil {
        log.Printf("Warning: Could not resolve absolute path for data directory: %v", err)
        absDataPath = dataPath
    }
    
    return &Config{
        Port:     getEnv("PORT", "4000"),
        DataPath: absDataPath,
        // DBHost: getEnv("DB_HOST", "localhost"),
        // JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
    }
}

// Helper function to get environment variable with default
func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}