package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"product-review/database"
	"product-review/models"
)

// Lấy tất cả reviews
func GetReviews(c echo.Context) error {
	var reviews []models.Review
	database.DB.Find(&reviews)
	return c.JSON(http.StatusOK, reviews)
}

// Tạo review
func CreateReview(c echo.Context) error {
	var review models.Review
	if err := c.Bind(&review); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	database.DB.Create(&review)
	return c.JSON(http.StatusCreated, review)
}
