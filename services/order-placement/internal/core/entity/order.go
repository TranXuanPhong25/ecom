package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CustomBaseModel - Base model for all entities
type CustomBaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Order - Domain entity representing an order
type Order struct {
	CustomBaseModel

	UserID      uuid.UUID `gorm:"type:uuid;not null;index"`
	Status      string    `gorm:"type:varchar(50);not null;default:'processing';index"`
	TotalAmount float64   `gorm:"type:decimal(15,2);not null;default:0"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}

// OrderItem - Domain entity representing an order item
type OrderItem struct {
	CustomBaseModel

	OrderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index"`

	Price    float64 `gorm:"type:decimal(15,2);not null;default:0"`
	Quantity int     `gorm:"type:int;not null"`
}

// OrderItemInput - Input data for creating order items
type OrderItemInput struct {
	ProductID uuid.UUID `json:"product_id" validate:"required,uuid"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
}

// CreateOrderRequest - Request to create a new order
type CreateOrderRequest struct {
	UserID uuid.UUID        `json:"user_id" validate:"required,uuid"`
	Items  []OrderItemInput `json:"items" validate:"required,min=1,dive"`
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
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateOrderResponse - Response after creating an order
type CreateOrderResponse struct {
	Order   OrderDTO `json:"order"`
	Message string   `json:"message"`
}
