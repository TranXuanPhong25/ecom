package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/TranXuanPhong25/ecom/carts/configs"
	"github.com/TranXuanPhong25/ecom/carts/controllers"
	"github.com/TranXuanPhong25/ecom/carts/repositories"
	"github.com/TranXuanPhong25/ecom/carts/routes"
	"github.com/TranXuanPhong25/ecom/carts/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs.LoadEnv()
	repositories.InitDBConnection()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderCookie, "X-User-Id"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowCredentials: true,
	}))

	serviceConfigs := configs.LoadServiceConfig()
	repo := repositories.NewCartRepository()
	productService := services.NewProductService(serviceConfigs)
	var shopsOnce sync.Once
	shopsService := services.NewShopsService(serviceConfigs, &shopsOnce)
	service := services.NewCartService(repo, productService, shopsService)
	controllers := controllers.NewCartController(service)
	// Register routes
	routes.RegisterCartRoutes(e, controllers)
	routes.RegisterHealthRoute(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
	}()

	<-quit
	shopsService.CloseConnection()

}
