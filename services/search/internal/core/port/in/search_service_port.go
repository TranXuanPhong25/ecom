package in

import (
	"context"

	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/dto"
)

// SearchServicePort defines the input port for search operations
// This is the interface that the application core provides to the outside world
type SearchServicePort interface {
	// SearchProducts searches for products based on a keyword query
	SearchProducts(ctx context.Context, request dto.SearchRequest) (*dto.SearchResponse, error)
}
