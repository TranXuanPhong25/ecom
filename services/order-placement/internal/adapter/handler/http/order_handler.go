package http

import (
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/port"
	"github.com/TranXuanPhong25/ecom/services/order-placement/validators"
	"github.com/labstack/echo/v4"
)

// OrderHandler handles HTTP requests for order placement
type OrderHandler struct {
	service port.OrderService
}

// NewOrderHandler creates a new OrderHandler
func NewOrderHandler(service port.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

// RegisterRoutes registers all order routes
func (h *OrderHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/api/orders/placement", h.CreateOrder)
}

// CreateOrder handles order creation request
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	req := new(dto.CreateOrderRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": "Invalid request format: " + err.Error(),
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": err.Error(),
		})
	}

	userIDHeader := c.Request().Header.Get("X-User-Id")
	if userIDHeader == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	res, svcErr := h.service.CreateOrder(req, userIDHeader)
	if svcErr != nil {
		return c.JSON(svcErr.Code, svcErr)
	}
	return c.JSON(http.StatusCreated, res)
}
