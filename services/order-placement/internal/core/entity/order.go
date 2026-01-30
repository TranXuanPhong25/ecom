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
