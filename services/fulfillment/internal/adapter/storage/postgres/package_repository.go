package postgres

import (
	"context"

	"github.com/rengumin/fulfillment/internal/core/dto"
	"github.com/rengumin/fulfillment/internal/core/entity"
	"gorm.io/gorm"
)

type packageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(db *gorm.DB) *packageRepository {
	return &packageRepository{db: db}
}

func (r *packageRepository) Create(ctx context.Context, pkg *entity.FulfillmentPackage) error {
	return r.db.WithContext(ctx).Create(pkg).Error
}

func (r *packageRepository) Update(ctx context.Context, pkg *entity.FulfillmentPackage) error {
	return r.db.WithContext(ctx).Save(pkg).Error
}

func (r *packageRepository) FindByPackageNumber(ctx context.Context, packageNumber string) (*entity.FulfillmentPackage, error) {
	var pkg entity.FulfillmentPackage
	err := r.db.WithContext(ctx).Where("package_number = ?", packageNumber).First(&pkg).Error
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *packageRepository) FindByOrderID(ctx context.Context, orderID int64) (*entity.FulfillmentPackage, error) {
	var pkg entity.FulfillmentPackage
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(&pkg).Error
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *packageRepository) FindAll(ctx context.Context, query dto.ListPackagesQuery) ([]entity.FulfillmentPackage, int64, error) {
	var packages []entity.FulfillmentPackage
	var total int64
	
	db := r.db.WithContext(ctx).Model(&entity.FulfillmentPackage{})
	
	// Apply filters
	if query.ShopID != "" {
		db = db.Where("shop_id = ?", query.ShopID)
	}
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	if query.Zone != "" {
		db = db.Where("delivery_zone = ?", query.Zone)
	}
	
	// Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Apply pagination
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 20
	}
	
	offset := (query.Page - 1) * query.PageSize
	
	// Fetch results
	err := db.Order("created_at DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&packages).Error
	
	if err != nil {
		return nil, 0, err
	}
	
	return packages, total, nil
}
