package routes

import (
	"github.com/TranXuanPhong25/ecom/product-review/controllers"
)

func ReviewRoutes(e *echo.Echo) {
	e.GET("/reviews", controllers.GetReviews)
	e.POST("/reviews", controllers.CreateReview)
}
