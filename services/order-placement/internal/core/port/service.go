package port

import (
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/dto"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// OrderService - Interface for order business logic
type OrderService interface {
	// CreateOrder creates a new order with items
	CreateOrder(request *dto.CreateOrderRequest, userID uuid.UUID) (*dto.CreateOrderResponse, *echo.HTTPError)
}
