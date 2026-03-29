package configs

import (
	"log"
	"os"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig Config

func LoadEnv() {
	AppConfig.ServerPort = getEnv("SERVER_PORT", "8080")
	AppConfig.DBHost = getEnv("DB_HOST", "localhost")
	AppConfig.DBPort = getEnv("DB_PORT", "5432")
	AppConfig.DBUser = getEnv("DB_USER", "postgres")
	AppConfig.DBPassword = getEnv("DB_PASSWORD", "postgres")
	AppConfig.DBName = getEnv("DB_NAME", "payment")
	log.Printf("Config loaded - Server Port: %s, DB_HOST: %s", AppConfig.ServerPort, AppConfig.DBHost)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
