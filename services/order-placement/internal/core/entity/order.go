package entity

import (
	"database/sql/driver"
	"encoding/json"
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

// OrderItem - Item structure stored as JSONB
type OrderItem struct {
	ProductID uuid.UUID `json:"producId"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
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

// Order - Domain entity representing an order with JSONB items
type Order struct {
	CustomBaseModel

	UserID      uuid.UUID  `gorm:"type:uuid;not null;index"`
	Status      string     `gorm:"type:varchar(50);not null;default:'processing';index"`
	TotalAmount float64    `gorm:"type:decimal(15,2);not null;default:0"`
	Items       OrderItems `gorm:"type:jsonb;not null;default:'[]'"`
}
