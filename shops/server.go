package main

import (
	"github.com/TranXuanPhong25/ecom/shops/middlewares"
	"github.com/TranXuanPhong25/ecom/shops/models"
	"github.com/TranXuanPhong25/ecom/shops/repositories"
	"github.com/TranXuanPhong25/ecom/shops/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.PrometheusMiddleware())

	repositories.ConnectDB()
	err := repositories.DB.AutoMigrate(&models.Shop{})
	if err != nil {
		e.Logger.Fatal("Failed to migrate database:", err)
	}
	routes.ShopsRoute(e)
	e.Logger.Fatal(e.Start(":8080"))

}
