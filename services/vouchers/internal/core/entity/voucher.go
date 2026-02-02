package entity

import "time"

// DiscountType represents the type of discount a voucher provides
type DiscountType string

const (
	DiscountTypePercentage  DiscountType = "PERCENTAGE"
	DiscountTypeFixedAmount DiscountType = "FIXED_AMOUNT"
)

// Voucher represents a discount voucher entity
type Voucher struct {
	ID            uint         `gorm:"primaryKey"`
	Code          string       `gorm:"uniqueIndex;not null;size:50"`
	DiscountType  DiscountType `gorm:"type:varchar(20);not null"`
	DiscountValue float64      `gorm:"not null"`
	MinOrderValue float64      `gorm:"default:0"`
	MaxUsage      int          `gorm:"default:0"` // 0 = unlimited (for future use)
	UsedCount     int          `gorm:"default:0"` // For future use
	ExpiresAt     *time.Time
	IsActive      bool   `gorm:"default:true"`
	Description   string `gorm:"size:500"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
