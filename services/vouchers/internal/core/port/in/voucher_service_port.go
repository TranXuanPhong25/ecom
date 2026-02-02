package in

import (
	"context"

	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/entity"
)

// VoucherServicePort defines the input port for voucher service operations
type VoucherServicePort interface {
	// CreateVoucher creates a new voucher
	CreateVoucher(ctx context.Context, request dto.CreateVoucherRequest) (*entity.Voucher, error)

	// GetVoucherByCode retrieves a voucher by its code
	GetVoucherByCode(ctx context.Context, code string) (*entity.Voucher, error)

	// GetVoucherByID retrieves a voucher by its ID
	GetVoucherByID(ctx context.Context, id uint) (*entity.Voucher, error)

	// ListVouchers retrieves a list of vouchers with pagination
	ListVouchers(ctx context.Context, request dto.ListVouchersRequest) ([]entity.Voucher, int64, error)

	// UpdateVoucher updates an existing voucher
	UpdateVoucher(ctx context.Context, id uint, request dto.UpdateVoucherRequest) (*entity.Voucher, error)

	// DeleteVoucher deletes a voucher by its ID
	DeleteVoucher(ctx context.Context, id uint) error
}
