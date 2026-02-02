package dto

import (
	"time"

	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/entity"
)

// CreateVoucherRequest represents the request to create a new voucher
type CreateVoucherRequest struct {
	Code          string               `json:"code" validate:"required,min=4,max=50"`
	DiscountType  entity.DiscountType  `json:"discount_type" validate:"required,oneof=PERCENTAGE FIXED_AMOUNT"`
	DiscountValue float64              `json:"discount_value" validate:"required,gt=0"`
	MinOrderValue float64              `json:"min_order_value" validate:"omitempty,gte=0"`
	MaxUsage      int                  `json:"max_usage" validate:"omitempty,gte=0"`
	ExpiresAt     *time.Time           `json:"expires_at" validate:"omitempty"`
	Description   string               `json:"description" validate:"omitempty,max=500"`
}

// UpdateVoucherRequest represents the request to update a voucher
type UpdateVoucherRequest struct {
	DiscountValue *float64   `json:"discount_value" validate:"omitempty,gt=0"`
	MinOrderValue *float64   `json:"min_order_value" validate:"omitempty,gte=0"`
	MaxUsage      *int       `json:"max_usage" validate:"omitempty,gte=0"`
	ExpiresAt     *time.Time `json:"expires_at" validate:"omitempty"`
	IsActive      *bool      `json:"is_active" validate:"omitempty"`
	Description   *string    `json:"description" validate:"omitempty,max=500"`
}

// VoucherResponse represents the response containing voucher data
type VoucherResponse struct {
	ID            uint                `json:"id"`
	Code          string              `json:"code"`
	DiscountType  entity.DiscountType `json:"discount_type"`
	DiscountValue float64             `json:"discount_value"`
	MinOrderValue float64             `json:"min_order_value"`
	MaxUsage      int                 `json:"max_usage"`
	UsedCount     int                 `json:"used_count"`
	ExpiresAt     *time.Time          `json:"expires_at"`
	IsActive      bool                `json:"is_active"`
	Description   string              `json:"description"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
}

// ListVouchersRequest represents the request to list vouchers with pagination
type ListVouchersRequest struct {
	Page   int   `query:"page" validate:"omitempty,gte=1"`
	Limit  int   `query:"limit" validate:"omitempty,gte=1,lte=100"`
	Active *bool `query:"active" validate:"omitempty"`
}

// ListVouchersResponse represents the response containing a list of vouchers
type ListVouchersResponse struct {
	Vouchers []VoucherResponse `json:"vouchers"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
}

// ToVoucherResponse converts entity.Voucher to VoucherResponse
func ToVoucherResponse(v *entity.Voucher) VoucherResponse {
	return VoucherResponse{
		ID:            v.ID,
		Code:          v.Code,
		DiscountType:  v.DiscountType,
		DiscountValue: v.DiscountValue,
		MinOrderValue: v.MinOrderValue,
		MaxUsage:      v.MaxUsage,
		UsedCount:     v.UsedCount,
		ExpiresAt:     v.ExpiresAt,
		IsActive:      v.IsActive,
		Description:   v.Description,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
	}
}
