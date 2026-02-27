package port

import (
	"context"
	"github.com/rengumin/fulfillment/internal/core/dto"
	"github.com/rengumin/fulfillment/internal/core/entity"
)

// PackageRepository defines data access for packages
type PackageRepository interface {
	Create(ctx context.Context, pkg *entity.FulfillmentPackage) error
	Update(ctx context.Context, pkg *entity.FulfillmentPackage) error
	FindByPackageNumber(ctx context.Context, packageNumber string) (*entity.FulfillmentPackage, error)
	FindByOrderID(ctx context.Context, orderID int64) (*entity.FulfillmentPackage, error)
	FindAll(ctx context.Context, query dto.ListPackagesQuery) ([]entity.FulfillmentPackage, int64, error)
}
