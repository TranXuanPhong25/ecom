package routes

import (
	"github.com/TranXuanPhong25/ecom/carts/configs"
	"github.com/TranXuanPhong25/ecom/carts/controllers"
	"github.com/TranXuanPhong25/ecom/carts/repositories"
	"github.com/TranXuanPhong25/ecom/carts/services"
	"github.com/labstack/echo/v4"
)

func RegisterCartRoutes(e *echo.Echo) {
	service_configs := configs.LoadServiceConfig()
	repo := repositories.NewCartRepository()
	productService := services.NewProductService(service_configs)
	service := services.NewCartService(repo, productService)
	controllers := controllers.NewCartController(service)
	e.GET("/api/carts/mine", controllers.GetCart)
	e.GET("/api/carts/mine/summary", controllers.GetCart)

	e.POST("/api/carts/mine/items", controllers.AddItemToCart)
	e.PUT("/api/carts/mine/items", controllers.UpdateCartItem)
	e.DELETE("/api/carts/mine/items", controllers.DeleteItemInCart)
}
