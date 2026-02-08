package routes

import (
	"github.com/TranXuanPhong25/ecom/services/product-reviews/controllers"
	"github.com/labstack/echo/v4"
)

func ReviewRoutes(e *echo.Echo) {
	ctrl := controllers.NewReviewController()

	// Review CRUD
	e.POST("/api/reviews", ctrl.CreateReview)
	e.GET("/api/reviews/:id", ctrl.GetReview)
	e.PUT("/api/reviews/:id", ctrl.UpdateReview)
	e.DELETE("/api/reviews/:id", ctrl.DeleteReview)

	// Product reviews
	e.GET("/api/reviews/products", ctrl.GetProductReviews)
	e.GET("/api/reviews/products/stats", ctrl.GetProductStats)

	// User reviews
	e.GET("/api/reviews/users", ctrl.GetUserReviews)
}
