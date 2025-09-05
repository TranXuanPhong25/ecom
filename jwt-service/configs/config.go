package configs

import (
	"encoding/base64"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	RpcPort   string
	SecretKey []byte
}

var AppConfig Config

var (
	SigningMethod = jwt.SigningMethodHS256
	ExpireTime    = time.Duration(36)
)

func LoadEnv() {
	AppConfig.RpcPort = getEnv("RPC_PORT", ":50050")

	secretKeyBase64 := os.Getenv("JWT_SECRET_KEY")
	if secretKeyBase64 == "" {
		log.Fatal("JWT_SECRET_KEY environment variable is not set")
	}
	decodedSecretKey, err := base64.StdEncoding.DecodeString(secretKeyBase64)
	if err != nil {
		log.Fatalf("Failed to decode JWT secret key: %v", err)
	}
	AppConfig.SecretKey = decodedSecretKey

	log.Printf("JWT Config loaded - RPC Port: %s", AppConfig.RpcPort)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
