package configs

import (
	"log"
	"os"
)

type Config struct {
	ServerPort  string
	ScyllaNodes string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
}

var AppConfig Config

func LoadEnv() {
	AppConfig.ServerPort = getEnv("SERVER_PORT", "8080")
	AppConfig.ScyllaNodes = getEnv("SCYLLA_NODES", "localhost")
	AppConfig.DBPassword = getEnv("DB_PASSWORD", "postgres")
	AppConfig.DBPort = getEnv("DB_PORT", "5432")
	AppConfig.DBUser = getEnv("DB_USER", "postgres")
	AppConfig.DBName = getEnv("DB_NAME", "mydatabase")
	AppConfig.DBHost = getEnv("DB_HOST", "localhost")
	log.Printf("Shops Config loaded - Server Port: %s, SCYLLA_NODES: %s",
		AppConfig.ServerPort, AppConfig.ScyllaNodes)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
