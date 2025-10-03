package controllers

import (
	"net/http"

	"github.com/TranXuanPhong25/ecom/carts/dtos"
	"github.com/TranXuanPhong25/ecom/carts/services"
	"github.com/TranXuanPhong25/ecom/carts/utils"
	"github.com/labstack/echo/v4"
)

type cartController struct {
	cartService services.ICartService
}

func NewCartController(cartService services.ICartService) *cartController {
	return &cartController{
		cartService: cartService,
	}
}

func (controller *cartController) AddItemToCart(c echo.Context) error {
	req := new(dtos.CartItemPayload)
	err := utils.ValidateRequestStructure(c, req)
	if err != nil {
		return err
	}
	userID := c.Request().Header["X-User-Id"][0]

	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	if err := controller.cartService.AddItemToCart(userID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item added to cart"})
}

func (controller *cartController) UpdateCartItem(c echo.Context) error {
	req := new(dtos.CartItemPayload)
	err := utils.ValidateRequestStructure(c, req)
	if err != nil {
		return err
	}
	userID := c.Request().Header["X-User-Id"][0]
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	if err := controller.cartService.UpdateItemInCart(userID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cart item updated"})
}

func (controller *cartController) DeleteItemInCart(c echo.Context) error {
	req := new(dtos.CartItemPayload)
	err := utils.ValidateRequestStructure(c, req)
	if err != nil {
		return err
	}
	userID := c.Request().Header["X-User-Id"][0]

	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	if err := controller.cartService.DeleteItemInCart(userID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (controller *cartController) GetCart(c echo.Context) error {
	userID := c.Request().Header["X-User-Id"][0]
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	cart, err := controller.cartService.GetCart(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cart)
}
