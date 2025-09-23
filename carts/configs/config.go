package configs

import (
	"log"
	"os"
)

type Config struct {
	ServerPort  string
	ScyllaNodes string
}

var AppConfig Config

func LoadEnv() {
	AppConfig.ServerPort = getEnv("SERVER_PORT", "8080")
	AppConfig.ScyllaNodes = getEnv("SCYLLA_NODES", "localhost")

	log.Printf("Shops Config loaded - Server Port: %s, SCYLLA_NODES: %s",
		AppConfig.ServerPort, AppConfig.ScyllaNodes)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
