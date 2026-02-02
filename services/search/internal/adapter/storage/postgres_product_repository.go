package storage

import (
	"context"
	"fmt"

	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/entity"
	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/port/out"
	"gorm.io/gorm"
)

// PostgresProductRepository implements the ProductRepositoryPort interface
// This is the adapter for PostgreSQL database access
type PostgresProductRepository struct {
	db *gorm.DB
}

// NewPostgresProductRepository creates a new PostgreSQL repository
func NewPostgresProductRepository(db *gorm.DB) out.ProductRepositoryPort {
	return &PostgresProductRepository{
		db: db,
	}
}

// SearchByKeyword performs full-text search on products
func (r *PostgresProductRepository) SearchByKeyword(ctx context.Context, keyword string, limit, offset int) ([]entity.Product, int64, error) {
	var products []entity.Product
	var total int64

	// PostgreSQL full-text search query using tsvector
	searchQuery := `
		SELECT id, name, description, category_id, cover_image, created_at, updated_at
		FROM products
		WHERE to_tsvector('english', name || ' ' || COALESCE(description, '')) @@ plainto_tsquery('english', ?)
		ORDER BY ts_rank(to_tsvector('english', name || ' ' || COALESCE(description, '')), plainto_tsquery('english', ?)) DESC, id ASC
		LIMIT ? OFFSET ?
	`

	// Count query
	countQuery := `
		SELECT COUNT(*)
		FROM products
		WHERE to_tsvector('english', name || ' ' || COALESCE(description, '')) @@ plainto_tsquery('english', ?)
	`

	// Execute count query
	if err := r.db.WithContext(ctx).Raw(countQuery, keyword).Scan(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	// Execute search query
	if err := r.db.WithContext(ctx).Raw(searchQuery, keyword, keyword, limit, offset).Scan(&products).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to search products: %w", err)
	}

	return products, total, nil
}
