package controllers

import (
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/carts/models"
	"github.com/TranXuanPhong25/ecom/carts/services"
	"github.com/TranXuanPhong25/ecom/carts/validators"
	"github.com/labstack/echo/v4"
)

func AddItemToCart(c echo.Context) error {
	req := new(models.CartItem)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": "Invalid request format",
		})
	}
	userID := c.QueryParam("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": err.Error(),
		})
	}
	//if err := services.AddItemToCart(userID, &item); err != nil {
	//	return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	//}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item added to cart"})
}

func GetCart(c echo.Context) error {

	userID := c.QueryParam("userId")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	cart, err := services.GetCart(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cart)
}
