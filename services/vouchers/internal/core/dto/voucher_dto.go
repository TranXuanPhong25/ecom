package dto

import (
	"time"

	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/entity"
)

// CreateVoucherRequest represents the request to create a new voucher
type CreateVoucherRequest struct {
	Code          string              `json:"code" validate:"required,min=4,max=50"`
	DiscountType  entity.DiscountType `json:"discountType" validate:"required,oneof=PERCENTAGE FIXED_AMOUNT"`
	DiscountValue float64             `json:"discountValue" validate:"required,gt=0"`
	MinOrderValue float64             `json:"minOrderValue" validate:"omitempty,gte=0"`
	MaxUsage      int                 `json:"maxUsage" validate:"omitempty,gte=0"`
	ExpiresAt     *time.Time          `json:"expiresAt" validate:"omitempty"`
	Description   string              `json:"description" validate:"omitempty,max=500"`
}

// UpdateVoucherRequest represents the request to update a voucher
type UpdateVoucherRequest struct {
	DiscountValue *float64   `json:"discountValue" validate:"omitempty,gt=0"`
	MinOrderValue *float64   `json:"minOrderValue" validate:"omitempty,gte=0"`
	MaxUsage      *int       `json:"maxUsage" validate:"omitempty,gte=0"`
	ExpiresAt     *time.Time `json:"expiresAt" validate:"omitempty"`
	IsActive      *bool      `json:"isActive" validate:"omitempty"`
	Description   *string    `json:"description" validate:"omitempty,max=500"`
}

// VoucherResponse represents the response containing voucher data
type VoucherResponse struct {
	ID            uint                `json:"id"`
	Code          string              `json:"code"`
	DiscountType  entity.DiscountType `json:"discountType"`
	DiscountValue float64             `json:"discountValue"`
	MinOrderValue float64             `json:"minOrderValue"`
	MaxUsage      int                 `json:"maxUsage"`
	UsedCount     int                 `json:"usedCount"`
	ExpiresAt     *time.Time          `json:"expiresAt"`
	IsActive      bool                `json:"isActive"`
	Description   string              `json:"description"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
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
