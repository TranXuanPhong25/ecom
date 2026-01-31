package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	httpHandler "github.com/TranXuanPhong25/ecom/services/order-placement/internal/adapter/handler/http"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/adapter/storage/postgres"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/config"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/entity"
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
	e.Use(middleware.CORS())

	// Register routes
	handler.RegisterRoutes(e)

	return e
}

func main() {
	// Load configuration
	config.LoadEnv()

	// Connect to database
	db := postgres.ConnectDB()

	// Auto migrate database schemas
	err := db.AutoMigrate(
		&entity.Order{},
	)
	if err != nil {
		fmt.Printf("Failed to migrate database: %v\n", err)
	}

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
