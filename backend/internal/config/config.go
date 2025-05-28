package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    Port string
    // Add other config fields as needed
    // DBHost string
    // JWTSecret string
}

func Load() *Config {
    // Try to load .env file (don't fail if it doesn't exist)
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Println("No .env file found, using environment variables and defaults")
    }
    
    return &Config{
        Port: getEnv("PORT", "8080"),
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