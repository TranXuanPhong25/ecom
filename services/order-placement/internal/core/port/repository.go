package port

import (
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/entity"
)

// OrderRepository - Interface for order data access
type OrderRepository interface {
	// CreateOrderWithItems creates an order with its items in a transaction
	CreateOrderWithItems(order *entity.Order, items []dto.OrderItemInput) (*entity.Order, error)
}
