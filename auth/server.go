package main

import (
	"github.com/TranXuanPhong25/ecom/auth/repositories"
	"github.com/TranXuanPhong25/ecom/auth/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	repositories.ConnectRedis()
	repositories.TestRedis()

	routes.AuthRoute(e)

	e.Logger.Fatal(e.Start(":8202"))
}
