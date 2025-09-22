package main

import (
	"log"

	"github.com/TranXuanPhong25/ecom/carts/configs"
	"github.com/TranXuanPhong25/ecom/carts/repositories"
	"github.com/TranXuanPhong25/ecom/carts/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	repositories.InitDBConnection()

	e := echo.New()

	routes.RegisterCartRoutes(e)
	routes.RegisterHealthRoute(e)
	// Start server
	if err := e.Start(":" + configs.AppConfig.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
