package config

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
	AppConfig.ServerPort = getEnv("SERVER_PORT", ":8080")
	AppConfig.DBHost = getEnv("DB_HOST", "localhost")
	AppConfig.DBPort = getEnv("DB_PORT", "5432")
	AppConfig.DBUser = getEnv("DB_USER", "postgres")
	AppConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	AppConfig.DBName = getEnv("DB_NAME", "shops_db")

	log.Printf("Shops Config loaded - Server Port: %s, DB: %s:%s",
		AppConfig.ServerPort, AppConfig.DBHost, AppConfig.DBPort)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
