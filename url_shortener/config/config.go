package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DB_URL    string
	RedisAddr string
}

// LoadConfig loads environment variables from `.env` file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using default values.")
	}

	return &Config{
		Port:      getEnv("PORT", "8080"),
		DB_URL:    getEnv("DATABASE_URL", "postgres://postgres:quocdat2602@localhost:5432/DATABASE_URL?sslmode=disable"),
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

// Helper function to get env variable or default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
