package routes

import (
	"github.com/TranXuanPhong25/ecom/services/product-reviews/controllers"
	"github.com/labstack/echo/v4"
)

func ReviewRoutes(e *echo.Echo) {
	ctrl := controllers.NewReviewController()

	// Review CRUD
	e.POST("/reviews", ctrl.CreateReview)
	e.GET("/reviews/:id", ctrl.GetReview)
	e.PUT("/reviews/:id", ctrl.UpdateReview)
	e.DELETE("/reviews/:id", ctrl.DeleteReview)

	// Product reviews
	e.GET("/products/:productId/reviews", ctrl.GetProductReviews)
	e.GET("/products/:productId/reviews/stats", ctrl.GetProductStats)

	// User reviews
	e.GET("/users/:userId/reviews", ctrl.GetUserReviews)
}

