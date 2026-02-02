package out

import (
	"context"

	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/entity"
)

// ProductRepositoryPort defines the output port for product data access
// This is the interface that the application core needs from external systems
type ProductRepositoryPort interface {
	// SearchByKeyword searches for products using full-text search
	SearchByKeyword(ctx context.Context, keyword string, limit, offset int) ([]entity.Product, int64, error)
}
