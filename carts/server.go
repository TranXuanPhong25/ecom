package main

import (
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
	e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
}
