package main

import (
	"github.com/TranXuanPhong25/ecom/carts/configs"
	"github.com/TranXuanPhong25/ecom/carts/repositories"
	"github.com/TranXuanPhong25/ecom/carts/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs.LoadEnv()
	repositories.InitDBConnection()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderCookie, "X-User-Id"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowCredentials: true,
	}))

	// Register routes
	routes.RegisterCartRoutes(e)
	routes.RegisterHealthRoute(e)
	// Start server
	e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
}
