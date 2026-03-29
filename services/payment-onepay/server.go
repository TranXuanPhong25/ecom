package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TranXuanPhong25/ecom/services/payment-onepay/configs"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/controllers"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/repositories"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/routes"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/services"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	repositories.InitDBConnection()

	e := echo.New()

	onepayConfig := configs.LoadServiceConfig()
	repo := repositories.NewPaymentRepository()
	onepayService := services.NewOnepayService(onepayConfig)
	service := services.NewPaymentService(repo, onepayService)
	paymentController := controllers.NewPaymentController(service)

	routes.RegisterPaymentRoutes(e, paymentController)
	routes.RegisterHealthRoute(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
	}()

	<-quit
}
