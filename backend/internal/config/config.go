package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DataPath  string
	JWTSecret string
	// Add other config fields as needed
	// DBHost string
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
	_ = godotenv.Load("../../.env") // Still optional

	// Try to walk upward to find the .git folder or backend folder (project root)
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory")
	}

	// Traverse until we find the project root
	projectRoot := cwd
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(filepath.Join(projectRoot, "backend", "data")); err == nil {
			break
		}
		projectRoot = filepath.Dir(projectRoot)
	}

	// Default fallback
	dataPath := filepath.Join(projectRoot, "backend", "data")
	if override := os.Getenv("DATA_PATH"); override != "" {
		dataPath = override
	}

	return &Config{
		Port:      getEnv("PORT", "4000"),
		DataPath:  dataPath,
		JWTSecret: getEnv("JWT_SECRET", "your-256-bit-secret"),
	}
}

// Helper function to get environment variable with default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}