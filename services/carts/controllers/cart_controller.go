package controllers

import (
	"net/http"

	"github.com/TranXuanPhong25/ecom/services/carts/dtos"
	"github.com/TranXuanPhong25/ecom/services/carts/services"
	"github.com/TranXuanPhong25/ecom/services/carts/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ICartController interface {
	AddItemToCart(c echo.Context) error
	UpdateCartItem(c echo.Context) error
	DeleteItemInCart(c echo.Context) error
	GetCart(c echo.Context) error
	GetTotalItemsInCart(c echo.Context) error
}
type CartController struct {
	cartService services.ICartService
}

func NewCartController(cartService services.ICartService) ICartController {
	return &CartController{
		cartService: cartService,
	}
}

func (controller *CartController) AddItemToCart(c echo.Context) error {
	req := new(dtos.CartItemPayload)
	err := utils.ValidateRequestStructure(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	userID := c.Request().Header["X-User-Id"][0]

	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	safeUserID, err := uuid.Parse(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid User ID"})
	}
	if err := controller.cartService.AddItemToCart(safeUserID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item added to cart"})
}

func (controller *CartController) UpdateCartItem(c echo.Context) error {
	req := new(dtos.CartItemPayload)
	err := utils.ValidateRequestStructure(c, req)
	if err != nil {
		return err
	}
	userID := c.Request().Header["X-User-Id"][0]
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	safeUserID, err := uuid.Parse(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid User ID"})
	}
	if err := controller.cartService.UpdateItemInCart(safeUserID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cart item updated"})
}

func (controller *CartController) DeleteItemInCart(c echo.Context) error {
	//idsParam := c.QueryParam("ids")
	//uuids := strings.Split(idsParam, ",")
	//for _, u := range uuids {
	//	_, err := uuid.Parse(strings.TrimSpace(u))
	//	if err != nil {
	//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	//	}
	//}
	req := new(dtos.DeleteCartItemsPayload)
	err := utils.ValidateRequestStructure(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	uuids := req.Items
	userID := c.Request().Header.Get("X-User-Id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	if err := controller.cartService.DeleteItemInCart(userID, uuids); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (controller *CartController) GetCart(c echo.Context) error {
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

func (controller *CartController) GetTotalItemsInCart(c echo.Context) error {
	userID := c.Request().Header["X-User-Id"][0]
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	totalItems, err := controller.cartService.GetTotalItems(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int{"totalItems": totalItems})
}
