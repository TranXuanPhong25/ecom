package controllers

import (
	"net/http"

	"github.com/TranXuanPhong25/ecom/carts/dtos"
	"github.com/TranXuanPhong25/ecom/carts/services"
	"github.com/TranXuanPhong25/ecom/carts/utils"
	"github.com/labstack/echo/v4"
)

func AddItemToCart(c echo.Context) error {
	req := new(dtos.AddItemRequest)
	err := utils.ValidateRequestStructure(c, req)
	if err != nil {
		return err
	}
	userID := c.Request().Header["X-User-Id"][0]

	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	if err := services.AddItemToCart(userID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item added to cart"})
}

func UpdateCartItem(c echo.Context) error {
	req := new(dtos.AddItemRequest)
	err := utils.ValidateRequestStructure(c, req)
	if err != nil {
		return err
	}
	userID := c.Request().Header["X-User-Id"][0]
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	if err := services.UpdateItemInCart(userID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cart item updated"})
}

func DeleteItemInCart(c echo.Context) error {
	req := new(dtos.AddItemRequest)
	err := utils.ValidateRequestStructure(c, req)
	if err != nil {
		return err
	}
	userID := c.Request().Header["X-User-Id"][0]

	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	if err := services.DeleteItemInCart(userID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func GetCart(c echo.Context) error {
	userID := c.Request().Header["X-User-Id"][0]
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	cart, err := services.GetCart(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cart)
}
