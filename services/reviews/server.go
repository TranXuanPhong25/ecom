package main

import (
	"github.com/TranXuanPhong25/ecom/services/product-reviews/configs"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/database"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/models"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configuration
	configs.LoadEnv()

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to database
	database.ConnectDB()

	// Auto migrate tables
	err := database.DB.AutoMigrate(&models.Review{})
	if err != nil {
		e.Logger.Fatal("Failed to migrate database:", err)
		return
	}

	// Register routes
	routes.ReviewRoutes(e)

	// Start server
	serverPort := configs.AppConfig.ServerPort
	e.Logger.Fatal(e.Start(serverPort))
}
