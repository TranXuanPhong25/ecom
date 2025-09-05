package config

import (
	"log"
	"os"
)

type Config struct {
	ServerPort      string
	MinIOEndpoint   string
	MinIOAccessKey  string
	MinIOSecretKey  string
	MinIOBucketName string
	MinIOUseSSL     string
}

var AppConfig Config

func LoadEnv() {
	AppConfig.ServerPort = getEnv("SERVER_PORT", ":8080")
	AppConfig.MinIOEndpoint = getEnv("MINIO_ENDPOINT", "localhost:9000")
	AppConfig.MinIOAccessKey = getEnv("MINIO_ACCESS_KEY", "minioadmin")
	AppConfig.MinIOSecretKey = getEnv("MINIO_SECRET_KEY", "minioadmin")
	AppConfig.MinIOBucketName = getEnv("MINIO_BUCKET_NAME", "uploads")
	AppConfig.MinIOUseSSL = getEnv("MINIO_USE_SSL", "false")

	log.Printf("Upload Config loaded - Server Port: %s, MinIO: %s",
		AppConfig.ServerPort, AppConfig.MinIOEndpoint)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
