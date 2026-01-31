package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// OrderItem - Item structure stored as JSONB
type OrderItem struct {
	ProductID     string `json:"productId"`
	ProductName   string `json:"productName"`
	ProductSku    string `json:"productSku,omitempty"`
	ImageUrl      string `json:"imageUrl"`
	VariantID     string `json:"variantId,omitempty"`
	VariantName   string `json:"variantName,omitempty"`
	OriginalPrice int64  `json:"originalPrice"`
	SalePrice     int64  `json:"salePrice"`
	Quantity      int    `json:"quantity"`
}

// OrderItems - Custom type for JSONB array
type OrderItems []OrderItem

// Scan implements sql.Scanner interface for reading from database
func (oi *OrderItems) Scan(value interface{}) error {
	if value == nil {
		*oi = OrderItems{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, oi)
}

// Value implements driver.Valuer interface for writing to database
func (oi OrderItems) Value() (driver.Value, error) {
	if len(oi) == 0 {
		return json.Marshal([]OrderItem{})
	}
	return json.Marshal(oi)
}

// DiscountInfo - Custom type for JSONB discount
type DiscountInfo map[string]interface{}

// Scan implements sql.Scanner interface
func (d *DiscountInfo) Scan(value interface{}) error {
	if value == nil {
		*d = DiscountInfo{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, d)
}

// Value implements driver.Valuer interface
func (d DiscountInfo) Value() (driver.Value, error) {
	if d == nil {
		return json.Marshal(map[string]interface{}{})
	}
	return json.Marshal(d)
}

// Order - Domain entity representing an order
type Order struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`

	OrderNumber string `gorm:"type:varchar(50);uniqueIndex;not null"`
	UserID      string `gorm:"type:varchar(255);not null;index"`
	ShopID      string `gorm:"type:varchar(255);not null;index"`

	// Shipping information
	RecipientName   string `gorm:"type:varchar(255);not null"`
	RecipientPhone  string `gorm:"type:varchar(20);not null"`
	DeliveryAddress string `gorm:"type:text;not null"`

	// Order status
	Status string `gorm:"type:varchar(50);not null;default:'CREATED';index"`

	// Payment information
	PaymentMethod string     `gorm:"type:varchar(50);not null"`
	PaymentStatus string     `gorm:"type:varchar(50);not null;default:'UNPAID';index"`
	PaidAt        *time.Time `gorm:"type:timestamptz"`

	// Pricing
	Subtotal    int64        `gorm:"not null;default:0"`
	ShippingFee int64        `gorm:"not null;default:0"`
	Discount    DiscountInfo `gorm:"type:jsonb;default:'{}'"`
	TotalAmount int64        `gorm:"not null;default:0"`

	// Shipping
	ShippingMethod    string     `gorm:"type:varchar(50)"`
	ShippingProvider  string     `gorm:"type:varchar(100)"`
	TrackingNumber    string     `gorm:"type:varchar(100);index"`
	EstimatedDelivery *time.Time `gorm:"type:timestamptz"`
	ActualDelivery    *time.Time `gorm:"type:timestamptz"`

	// Notes
	CustomerNote string `gorm:"type:text"`
	SellerNote   string `gorm:"type:text"`
	CancelReason string `gorm:"type:text"`

	// Timestamps
	ConfirmedAt  *time.Time `gorm:"type:timestamptz"`
	ProcessingAt *time.Time `gorm:"type:timestamptz"`
	ShippedAt    *time.Time `gorm:"type:timestamptz"`
	DeliveredAt  *time.Time `gorm:"type:timestamptz"`
	CompletedAt  *time.Time `gorm:"type:timestamptz"`
	CancelledAt  *time.Time `gorm:"type:timestamptz"`
	CreatedAt    time.Time  `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt    time.Time  `gorm:"type:timestamptz;not null;default:now()"`

	// Items stored as JSONB
	Items OrderItems `gorm:"type:jsonb;not null;default:'[]'"`
}

// TableName sets the table name for Order
func (Order) TableName() string {
	return "orders"
}

// CalculateTotalAmount calculates subtotal and total amount
func (o *Order) CalculateTotalAmount() {
	var subtotal int64
	for _, item := range o.Items {
		subtotal += item.SalePrice * int64(item.Quantity)
	}
	o.Subtotal = subtotal
	o.TotalAmount = subtotal + o.ShippingFee
}
