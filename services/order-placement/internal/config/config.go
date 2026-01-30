package config

import (
	"log"
	"os"
)

// Config holds application configuration
type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// AppConfig is the global configuration instance
var AppConfig Config

// LoadEnv loads environment variables into configuration
func LoadEnv() {
	AppConfig.ServerPort = getEnv("SERVER_PORT", "8080")
	AppConfig.DBHost = getEnv("DB_HOST", "localhost")
	AppConfig.DBPort = getEnv("DB_PORT", "5432")
	AppConfig.DBUser = getEnv("DB_USER", "postgres")
	AppConfig.DBPassword = getEnv("DB_PASSWORD", "postgres")
	AppConfig.DBName = getEnv("DB_NAME", "mydatabase")

	log.Printf("Order Placement Config loaded - Server Port: %s, DB: %s:%s",
		AppConfig.ServerPort, AppConfig.DBHost, AppConfig.DBPort)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
