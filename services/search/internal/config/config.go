package config

import (
	"log"
	"os"
)

// Config holds all application configuration
type Config struct {
	ServerPort string
	Database   DatabaseConfig
}

// DatabaseConfig holds database connection configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	config := &Config{
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "mydatabase"),
		},
	}

	log.Printf("Configuration loaded - Server: %s, DB: %s:%s/%s",
		config.ServerPort,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName)

	return config
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
