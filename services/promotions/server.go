package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/TranXuanPhong25/ecom/services/promotions/configs"
	"github.com/TranXuanPhong25/ecom/services/promotions/middlewares"
	"github.com/TranXuanPhong25/ecom/services/promotions/models"
	"github.com/TranXuanPhong25/ecom/services/promotions/repositories"
	"github.com/TranXuanPhong25/ecom/services/promotions/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func newEchoServer() *echo.Echo {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middlewares.PrometheusMiddleware())
	routes.PromotionsRoute(e)
	return e
}

func main() {
	configs.LoadEnv()
	e := newEchoServer()
	repositories.ConnectDB()

	// Auto migrate database schemas
	err := repositories.DB.AutoMigrate(
		&models.EventBanner{},
		&models.PromoBar{},
	)
	if err != nil {
		fmt.Printf("Failed to migrate database: %v\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		log.Printf("HTTP server listening on %v\n", configs.AppConfig.ServerPort)
		e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutting down server...\n")
	e.Shutdown(context.Background())

	wg.Wait()
}
