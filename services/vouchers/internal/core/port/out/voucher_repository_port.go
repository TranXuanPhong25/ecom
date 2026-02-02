package out

import (
	"context"

	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/entity"
)

// VoucherRepositoryPort defines the output port for voucher repository operations
type VoucherRepositoryPort interface {
	// Create creates a new voucher in the repository
	Create(ctx context.Context, voucher *entity.Voucher) error

	// FindByCode retrieves a voucher by its code
	FindByCode(ctx context.Context, code string) (*entity.Voucher, error)

	// FindByID retrieves a voucher by its ID
	FindByID(ctx context.Context, id uint) (*entity.Voucher, error)

	// FindAll retrieves all vouchers with pagination and optional active filter
	FindAll(ctx context.Context, limit, offset int, active *bool) ([]entity.Voucher, int64, error)

	// Update updates an existing voucher in the repository
	Update(ctx context.Context, voucher *entity.Voucher) error

	// Delete deletes a voucher by its ID
	Delete(ctx context.Context, id uint) error
}
