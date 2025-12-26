package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/TranXuanPhong25/ecom/services/carts/configs"
	"github.com/TranXuanPhong25/ecom/services/carts/controllers"
	"github.com/TranXuanPhong25/ecom/services/carts/repositories"
	"github.com/TranXuanPhong25/ecom/services/carts/routes"
	"github.com/TranXuanPhong25/ecom/services/carts/services"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	repositories.InitDBConnection()

	e := echo.New()

	serviceConfigs := configs.LoadServiceConfig()
	repo := repositories.NewCartRepository()
	productService := services.NewProductService(serviceConfigs)
	var shopsOnce sync.Once
	shopsService := services.NewShopsService(serviceConfigs, &shopsOnce)
	service := services.NewCartService(repo, productService, shopsService)
	cartController := controllers.NewCartController(service)
	// Register routes
	routes.RegisterCartRoutes(e, cartController)
	routes.RegisterHealthRoute(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
	}()

	<-quit
	shopsService.CloseConnection()

}
