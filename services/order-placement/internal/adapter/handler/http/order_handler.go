package http

import (
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/entity"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/port"
	"github.com/TranXuanPhong25/ecom/services/order-placement/validators"
	"github.com/labstack/echo/v4"
)

// OrderHandler handles HTTP requests for orders
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
	e.POST("/api/orders", h.CreateOrder)
}

// CreateOrder handles order creation request
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	req := new(entity.CreateOrderRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": "Invalid request format",
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": err.Error(),
		})
	}

	res, err := h.service.CreateOrder(req)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusCreated, res)
}
