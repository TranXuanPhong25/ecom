package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	httpHandler "github.com/TranXuanPhong25/ecom/services/order-placement/internal/adapter/handler/http"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/adapter/storage/postgres"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/config"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func newEchoServer(handler *httpHandler.OrderHandler) *echo.Echo {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register routes
	handler.RegisterRoutes(e)

	return e
}

func runMigrations() {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	// Get migrations path from env or use default
	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		// Get current working directory and build absolute path
		cwd, err := os.Getwd()
		if err != nil {
			log.Printf("Failed to get working directory: %v", err)
			return
		}
		migrationsPath = "file://" + cwd + "/migrations"
	}

	log.Printf("Running migrations from: %s", migrationsPath)

	m, err := migrate.New(
		migrationsPath,
		dbURL,
	)
	if err != nil {
		log.Printf("Failed to create migrate instance: %v", err)
		return
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Failed to run migrations: %v", err)
		return
	}

	log.Print("Migrations completed successfully")
}

func main() {
	// Load configuration
	config.LoadEnv()

	// Run SQL migrations
	runMigrations()

	// Connect to database
	db := postgres.ConnectDB()

	// Initialize layers (Dependency Injection)
	orderRepo := postgres.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := httpHandler.NewOrderHandler(orderService)

	// Create Echo server
	e := newEchoServer(orderHandler)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("HTTP server listening on %v\n", config.AppConfig.ServerPort)
		if err := e.Start(":" + config.AppConfig.ServerPort); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutting down server...\n")
	e.Shutdown(context.Background())

	wg.Wait()
}
