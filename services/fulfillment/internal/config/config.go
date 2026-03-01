package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Kafka
	KafkaBrokers string
	TLSConfig    *tls.Config
	// Fulfillment settings
	DefaultPickupWindow   int // hours ahead to schedule pickup
	MaxDeliveryAttempts   int
	EstimatedDeliveryDays int
}

func LoadConfig() (*Config, error) {
	// Load .env file if exists
	godotenv.Load()

	kafkaTruststoreLocation := getEnv("KAFKA_SSL_CA_CERT_LOCATION", "ca.cert")
	caCert, err := os.ReadFile(kafkaTruststoreLocation)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded CA certificate from: %s", kafkaTruststoreLocation)
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Failed to append CA certificate to pool")
	}

	tlsConfig := &tls.Config{
		RootCAs:    caCertPool,
		MinVersion: tls.VersionTLS12,
	}

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "mydatabase"),

		KafkaBrokers: getEnv("KAFKA_BROKERS", "localhost:31092"),
		TLSConfig:    tlsConfig,

		DefaultPickupWindow:   getEnvInt("DEFAULT_PICKUP_WINDOW", 24),
		MaxDeliveryAttempts:   getEnvInt("MAX_DELIVERY_ATTEMPTS", 3),
		EstimatedDeliveryDays: getEnvInt("ESTIMATED_DELIVERY_DAYS", 3),
	}, nil
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	var result int
	fmt.Sscanf(value, "%d", &result)
	return result
}
