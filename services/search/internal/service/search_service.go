package service

import (
	"context"
	"errors"
	"strings"

	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/port/in"
	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/port/out"
)

// SearchService implements the SearchServicePort interface
// This is the application's business logic layer
type SearchService struct {
	productRepo out.ProductRepositoryPort
}

// NewSearchService creates a new SearchService with dependency injection
func NewSearchService(productRepo out.ProductRepositoryPort) in.SearchServicePort {
	return &SearchService{
		productRepo: productRepo,
	}
}

// SearchProducts implements the search use case
func (s *SearchService) SearchProducts(ctx context.Context, request dto.SearchRequest) (*dto.SearchResponse, error) {
	// Sanitize and validate query
	query := strings.TrimSpace(request.Query)
	if query == "" {
		return nil, errors.New("search query cannot be empty")
	}

	// Set default pagination values
	page := request.Page
	if page < 1 {
		page = 1
	}

	limit := request.Limit
	if limit < 1 {
		limit = 20 // default
	}
	if limit > 100 {
		limit = 100 // max
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Call repository to perform search
	products, total, err := s.productRepo.SearchByKeyword(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert entities to DTOs
	productDTOs := make([]dto.ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i] = dto.ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			CategoryID:  product.CategoryID,
			CoverImage:  product.CoverImage,
		}
	}

	return &dto.SearchResponse{
		Products: productDTOs,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}
