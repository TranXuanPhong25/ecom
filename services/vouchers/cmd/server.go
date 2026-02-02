package main

import (
	"log"

	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/adapter/handler"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/adapter/storage"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/config"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := config.InitDatabase(cfg.GetDSN())
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Wire dependencies (Hexagonal Architecture - Dependency Injection)
	// 1. Create repository (adapter OUT - implements output port)
	voucherRepository := storage.NewPostgresVoucherRepository(db)

	// 2. Create service (core - depends on repository port)
	voucherService := service.NewVoucherService(voucherRepository)

	// 3. Create handler (adapter IN - depends on service port)
	voucherHandler := handler.NewVoucherHandler(voucherService)

	// Initialize Echo server
	e := echo.New()

	// Register routes
	handler.RegisterRoutes(e, voucherHandler)

	// Start server
	serverAddr := ":" + cfg.ServerPort
	log.Printf("Starting voucher service on %s", serverAddr)
	if err := e.Start(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
