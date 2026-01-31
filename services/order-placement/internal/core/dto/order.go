package dto

import (
	"time"
)

// OrderItemInput - Input data for creating order items
type OrderItemInput struct {
	ProductID     string `json:"productId" validate:"required"`
	ProductName   string `json:"productName" validate:"required"`
	ProductSku    string `json:"productSku"`
	ImageUrl      string `json:"imageUrl" validate:"required"`
	VariantID     string `json:"variantId"`
	VariantName   string `json:"variantName"`
	OriginalPrice int64  `json:"originalPrice" validate:"required,min=0"`
	SalePrice     int64  `json:"salePrice" validate:"required,min=0"`
	Quantity      int    `json:"quantity" validate:"required,min=1"`
}

// CreateOrderRequest - Request to create a new order
type CreateOrderRequest struct {
	// Shop
	ShopID string `json:"shopId" validate:"required"`

	// Shipping information
	RecipientName   string `json:"recipientName" validate:"required"`
	RecipientPhone  string `json:"recipientPhone" validate:"required"`
	DeliveryAddress string `json:"deliveryAddress" validate:"required"`

	// Payment
	PaymentMethod string `json:"paymentMethod" validate:"required"`

	// Shipping
	ShippingMethod string `json:"shippingMethod"`
	ShippingFee    int64  `json:"shippingFee"`

	// Discount
	Discount map[string]interface{} `json:"discount"`

	// Notes
	CustomerNote string `json:"customerNote"`

	// Items
	Items []OrderItemInput `json:"items" validate:"required,min=1,dive"`
}

// OrderDTO - Data transfer object for order
type OrderDTO struct {
	ID          int64  `json:"id"`
	OrderNumber string `json:"orderNumber"`
	UserID      string `json:"userId"`
	ShopID      string `json:"shopId"`

	// Shipping information
	RecipientName   string `json:"recipientName"`
	RecipientPhone  string `json:"recipientPhone"`
	DeliveryAddress string `json:"deliveryAddress"`

	// Status
	Status        string     `json:"status"`
	PaymentMethod string     `json:"paymentMethod"`
	PaymentStatus string     `json:"paymentStatus"`
	PaidAt        *time.Time `json:"paidAt,omitempty"`

	// Pricing
	Subtotal    int64                  `json:"subtotal"`
	ShippingFee int64                  `json:"shippingFee"`
	Discount    map[string]interface{} `json:"discount"`
	TotalAmount int64                  `json:"totalAmount"`

	// Shipping details
	ShippingMethod    string     `json:"shippingMethod,omitempty"`
	ShippingProvider  string     `json:"shippingProvider,omitempty"`
	TrackingNumber    string     `json:"trackingNumber,omitempty"`
	EstimatedDelivery *time.Time `json:"estimatedDelivery,omitempty"`
	ActualDelivery    *time.Time `json:"actualDelivery,omitempty"`

	// Notes
	CustomerNote string `json:"customerNote,omitempty"`
	SellerNote   string `json:"sellerNote,omitempty"`
	CancelReason string `json:"cancelReason,omitempty"`

	// Timestamps
	ConfirmedAt  *time.Time `json:"confirmedAt,omitempty"`
	ProcessingAt *time.Time `json:"processingAt,omitempty"`
	ShippedAt    *time.Time `json:"shippedAt,omitempty"`
	DeliveredAt  *time.Time `json:"deliveredAt,omitempty"`
	CompletedAt  *time.Time `json:"completedAt,omitempty"`
	CancelledAt  *time.Time `json:"cancelledAt,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`

	OrderItems []OrderItemDTO `json:"orderItems"`
}

// OrderItemDTO - Data transfer object for order item
type OrderItemDTO struct {
	ProductID     string `json:"productId"`
	ProductName   string `json:"productName"`
	ProductSku    string `json:"productSku,omitempty"`
	ImageUrl      string `json:"imageUrl"`
	VariantID     string `json:"variantId,omitempty"`
	VariantName   string `json:"variantName,omitempty"`
	OriginalPrice int64  `json:"originalPrice"`
	SalePrice     int64  `json:"salePrice"`
	Quantity      int    `json:"quantity"`
	Subtotal      int64  `json:"subtotal"`
}

// CreateOrderResponse - Response after creating an order
type CreateOrderResponse struct {
	Order   OrderDTO `json:"order"`
	Message string   `json:"message"`
}
