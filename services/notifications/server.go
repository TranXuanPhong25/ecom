package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TranXuanPhong25/ecom/services/notifications/configs"
	"github.com/TranXuanPhong25/ecom/services/notifications/controllers"
	"github.com/TranXuanPhong25/ecom/services/notifications/repositories"
	"github.com/TranXuanPhong25/ecom/services/notifications/routes"
	"github.com/TranXuanPhong25/ecom/services/notifications/services"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	repositories.InitDBConnection()

	e := echo.New()

	repo := repositories.NewNotificationRepository()
	service := services.NewNotificationService(repo)
	notificationController := controllers.NewNotificationController(service)

	routes.RegisterNotificationRoutes(e, notificationController)
	routes.RegisterHealthRoute(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
	}()

	<-quit
}
