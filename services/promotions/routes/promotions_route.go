package routes

import (
	"github.com/TranXuanPhong25/ecom/services/promotions/controllers"
	"github.com/labstack/echo/v4"
)

func PromotionsRoute(e *echo.Echo) {
	// Event Banner routes
	e.POST("/api/promotions/banners", controllers.CreateEventBanner)
	e.GET("/api/promotions/banners", controllers.GetAllEventBanners)
	e.GET("/api/promotions/banners/:id", controllers.GetEventBannerByID)
	e.GET("/api/promotions/banners/type/:type", controllers.GetEventBannersByType)
	e.GET("/api/promotions/banners/position/:position", controllers.GetEventBannersByPosition)
	e.PUT("/api/promotions/banners", controllers.UpdateEventBanner)
	e.DELETE("/api/promotions/banners/:id", controllers.DeleteEventBanner)

	// Promo Bar routes
	e.POST("/api/promotions/bars", controllers.CreatePromoBar)
	e.GET("/api/promotions/bars", controllers.GetAllPromoBars)
	e.GET("/api/promotions/bars/:id", controllers.GetPromoBarByID)
	e.PUT("/api/promotions/bars", controllers.UpdatePromoBar)
	e.DELETE("/api/promotions/bars/:id", controllers.DeletePromoBar)

	// Combined route - get all active promotions
	e.GET("/api/promotions/active", controllers.GetActivePromotions)
}
