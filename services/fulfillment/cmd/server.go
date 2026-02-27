package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rengumin/fulfillment/internal/adapter/event"
	httpHandler "github.com/rengumin/fulfillment/internal/adapter/handler/http"
	postgresRepo "github.com/rengumin/fulfillment/internal/adapter/storage/postgres"
	"github.com/rengumin/fulfillment/internal/config"
	"github.com/rengumin/fulfillment/internal/core/entity"
	"github.com/rengumin/fulfillment/internal/service"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func runMigrations(cfg *config.Config) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Printf("Failed to get working directory: %v", err)
			return
		}
		migrationsPath = "file://" + cwd + "/migrations"
	}

	log.Printf("Running migrations from: %s", migrationsPath)

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Printf("Failed to create migrate instance: %v", err)
		return
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Failed to run migrations: %v", err)
		return
	}

	log.Println("Migrations completed successfully")
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Run SQL migrations
	runMigrations(cfg)

	// Connect to database
	db, err := connectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	packageRepo := postgresRepo.NewPackageRepository(db)

	// Initialize event publisher
	kafkaBrokers := strings.Split(cfg.KafkaBrokers, ",")
	publisher := event.NewKafkaPublisher(kafkaBrokers)
	err = createTopicIfNotExists(kafkaBrokers, string(entity.TopicDelivered))
	err = createTopicIfNotExists(kafkaBrokers, string(entity.TopicDeliveryFailed))
	err = createTopicIfNotExists(kafkaBrokers, string(entity.TopicPickedUp))
	err = createTopicIfNotExists(kafkaBrokers, string(entity.TopicPickupScheduled))
	err = createTopicIfNotExists(kafkaBrokers, string(entity.TopicInTransit))
	err = createTopicIfNotExists(kafkaBrokers, string(entity.TopicOutForDelivery))
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer publisher.Close()

	// Initialize services
	fulfillmentSvc := service.NewFulfillmentService(packageRepo, publisher, cfg)

	// Initialize HTTP handler
	handler := httpHandler.NewFulfillmentHandler(fulfillmentSvc)

	// Setup Echo server
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/health"
		},
		LogMethod:  true,
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogError:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			reqID, _ := c.Get("request_id").(string)
			fmt.Printf(
				"[HTTP] %s %s %d %s err=%v reqId=%s\n",
				v.Method,
				v.URI,
				v.Status,
				v.Latency,
				v.Error,
				reqID,
			)
			return nil
		},
	}))

	e.Use(middleware.Recover())
	handler.RegisterRoutes(e)

	// 🚚 Start auto-delivery simulator (cron job)
	simulator := service.NewAutoDeliverySimulator(packageRepo, publisher)
	go startDeliverySimulator(simulator)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	if err := e.Start(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func connectDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDBConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}

func createTopicIfNotExists(brokers []string, topic string) error {
	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka: %w", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("failed to get Kafka controller: %w", err)
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", fmt.Sprintf("%s:%d", controller.Host, controller.Port))
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka controller: %w", err)
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil && !strings.Contains(err.Error(), "Topic with this name already exists") {
		return fmt.Errorf("failed to create Kafka topic: %w", err)
	}

	log.Printf("Kafka topic '%s' is ready", topic)
	return nil
}

// startDeliverySimulator chạy cron job mỗi 30 phút
func startDeliverySimulator(simulator *service.AutoDeliverySimulator) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	log.Println("🚀 Auto-delivery simulator started (runs every 30 minutes)")

	// Run immediately on startup
	ctx := context.Background()
	if err := simulator.Run(ctx); err != nil {
		log.Printf("Error running simulator: %v", err)
	}

	// Then run every 30 minutes
	for range ticker.C {
		if err := simulator.Run(ctx); err != nil {
			log.Printf("Error running simulator: %v", err)
		}
	}
}
