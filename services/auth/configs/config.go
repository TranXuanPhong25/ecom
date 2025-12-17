package configs

import (
	"log"
	"os"
)

type Config struct {
	ServerPort       string
	JWTServiceAddr   string
	UsersServiceAddr string
	RedisHost        string
	RedisPort        string
}

var AppConfig Config

func LoadEnv() {
	AppConfig.ServerPort = getEnv("SERVER_PORT", "8080")
	AppConfig.JWTServiceAddr = getEnv("JWT_SERVICE_ADDR", "jwt-service:50050")
	AppConfig.UsersServiceAddr = getEnv("USERS_SERVICE_ADDR", "users-service:50050")
	AppConfig.RedisHost = getEnv("REDIS_HOST", "localhost")
	AppConfig.RedisPort = getEnv("REDIS_PORT", "6379")

	log.Printf("Config loaded - Server: %s, JWT: %s, Users: %s",
		AppConfig.ServerPort, AppConfig.JWTServiceAddr, AppConfig.UsersServiceAddr)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
