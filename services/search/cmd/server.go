package main

import (
	"log"

	"github.com/TranXuanPhong25/ecom/services/search-service/internal/adapter/handler"
	"github.com/TranXuanPhong25/ecom/services/search-service/internal/adapter/storage"
	"github.com/TranXuanPhong25/ecom/services/search-service/internal/config"
	"github.com/TranXuanPhong25/ecom/services/search-service/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Starting Search Service...")

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db, err := config.ConnectDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize layers (Hexagonal Architecture - Dependency Injection)
	
	// 1. Adapter Layer (OUT) - Repository
	productRepo := storage.NewPostgresProductRepository(db)

	// 2. Service Layer (Core) - Business Logic
	searchService := service.NewSearchService(productRepo)

	// 3. Adapter Layer (IN) - HTTP Server
	e := echo.New()
	handler.SetupRoutes(e, searchService)

	// Start server
	log.Printf("Server starting on %s", cfg.ServerPort)
	if err := e.Start(cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
