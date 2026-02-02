package storage

import (
	"context"
	"errors"

	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/entity"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/port/out"
	"gorm.io/gorm"
)

// PostgresVoucherRepository implements VoucherRepositoryPort using PostgreSQL
type PostgresVoucherRepository struct {
	db *gorm.DB
}

// NewPostgresVoucherRepository creates a new instance of PostgresVoucherRepository
func NewPostgresVoucherRepository(db *gorm.DB) out.VoucherRepositoryPort {
	return &PostgresVoucherRepository{
		db: db,
	}
}

// Create creates a new voucher in the database
func (r *PostgresVoucherRepository) Create(ctx context.Context, voucher *entity.Voucher) error {
	result := r.db.WithContext(ctx).Create(voucher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindByCode retrieves a voucher by its code
func (r *PostgresVoucherRepository) FindByCode(ctx context.Context, code string) (*entity.Voucher, error) {
	var voucher entity.Voucher
	result := r.db.WithContext(ctx).Where("code = ?", code).First(&voucher)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("voucher not found")
		}
		return nil, result.Error
	}
	return &voucher, nil
}

// FindByID retrieves a voucher by its ID
func (r *PostgresVoucherRepository) FindByID(ctx context.Context, id uint) (*entity.Voucher, error) {
	var voucher entity.Voucher
	result := r.db.WithContext(ctx).First(&voucher, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("voucher not found")
		}
		return nil, result.Error
	}
	return &voucher, nil
}

// FindAll retrieves all vouchers with pagination and optional active filter
func (r *PostgresVoucherRepository) FindAll(ctx context.Context, limit, offset int, active *bool) ([]entity.Voucher, int64, error) {
	var vouchers []entity.Voucher
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Voucher{})

	// Apply active filter if provided
	if active != nil {
		query = query.Where("is_active = ?", *active)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	result := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&vouchers)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return vouchers, total, nil
}

// Update updates an existing voucher in the database
func (r *PostgresVoucherRepository) Update(ctx context.Context, voucher *entity.Voucher) error {
	result := r.db.WithContext(ctx).Save(voucher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete deletes a voucher by its ID
func (r *PostgresVoucherRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&entity.Voucher{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("voucher not found")
	}
	return nil
}
