package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TranXuanPhong25/ecom/auth/middlewares"
	"github.com/TranXuanPhong25/ecom/auth/repositories"
	"github.com/TranXuanPhong25/ecom/auth/routes"
	"github.com/TranXuanPhong25/ecom/auth/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.PrometheusMiddleware())
	repositories.ConnectRedis()

	services.InitJWTServiceClient("jwt-service:50050")
	services.InitUsersServiceClient("users-service:50050")

	routes.AuthRoute(e)
	routes.MetricRoute(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	e.Logger.Fatal(e.Start(":8080"))

	<-quit
	services.CloseUsersServiceConnection()
	services.CloseJWTServiceConnection()
	repositories.CloseRedisConnection()

}
