package dto

import (
	"time"

	"github.com/google/uuid"
)

// OrderItemInput - Input data for creating order items
type OrderItemInput struct {
	ProductID uuid.UUID `json:"product_id" validate:"required,uuid"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
}

// CreateOrderRequest - Request to create a new order
type CreateOrderRequest struct {
	Items []OrderItemInput `json:"items" validate:"required,min=1,dive"`
}

// OrderDTO - Data transfer object for order
type OrderDTO struct {
	ID          uuid.UUID      `json:"id"`
	UserID      uuid.UUID      `json:"user_id"`
	TotalAmount float64        `json:"total_amount"`
	Status      string         `json:"status"`
	OrderItems  []OrderItemDTO `json:"order_items"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// OrderItemDTO - Data transfer object for order item
type OrderItemDTO struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}

// CreateOrderResponse - Response after creating an order
type CreateOrderResponse struct {
	Order   OrderDTO `json:"order"`
	Message string   `json:"message"`
}
