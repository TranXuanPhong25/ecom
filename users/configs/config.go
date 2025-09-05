package configs

import (
	"log"
	"os"
)

type Config struct {
	RpcPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig Config

func LoadEnv() {
	AppConfig.RpcPort = getEnv("RPC_PORT", ":50050")
	AppConfig.DBHost = getEnv("DB_HOST", "localhost")
	AppConfig.DBPort = getEnv("DB_PORT", "5432")
	AppConfig.DBUser = getEnv("DB_USER", "postgres")
	AppConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	AppConfig.DBName = getEnv("DB_NAME", "users_db")

	log.Printf("Users Config loaded - RPC Port: %s, DB: %s:%s",
		AppConfig.RpcPort, AppConfig.DBHost, AppConfig.DBPort)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
