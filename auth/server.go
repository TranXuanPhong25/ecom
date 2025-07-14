package main

import (
	"os"
	"os/signal"
	"syscall"

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

	repositories.ConnectRedis()
	repositories.TestRedis()

	services.InitJWTServiceClient("jwt-service:50051")
	services.InitUserServiceClient("user:50052")

	routes.AuthRoute(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	e.Logger.Fatal(e.Start(":8202"))

	<-quit
	services.CloseUserServiceConnection()
	services.CloseJWTServiceConnection()
	repositories.CloseRedisConnection()

}
