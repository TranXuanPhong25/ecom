package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TranXuanPhong25/ecom/auth/configs"
	"github.com/TranXuanPhong25/ecom/auth/middlewares"
	"github.com/TranXuanPhong25/ecom/auth/repositories"
	"github.com/TranXuanPhong25/ecom/auth/routes"
	"github.com/TranXuanPhong25/ecom/auth/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs.LoadEnv()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.PrometheusMiddleware())
	repositories.ConnectRedis(configs.AppConfig.RedisHost)

	services.InitJWTServiceClient(configs.AppConfig.JWTServiceAddr)
	services.InitUsersServiceClient(configs.AppConfig.UsersServiceAddr)

	routes.AuthRoute(e)
	routes.MetricRoute(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))

	<-quit
	services.CloseUsersServiceConnection()
	services.CloseJWTServiceConnection()
	repositories.CloseRedisConnection()

}
