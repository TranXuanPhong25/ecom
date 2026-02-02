package validators

import (
	"errors"
	"strings"

	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/dto"
)

// ValidateSearchRequest validates a search request
func ValidateSearchRequest(req *dto.SearchRequest) error {
	// Validate query
	query := strings.TrimSpace(req.Query)
	if query == "" {
		return errors.New("search query cannot be empty")
	}
	if len(query) > 200 {
		return errors.New("search query is too long (max 200 characters)")
	}

	// Validate pagination
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 {
		req.Limit = 20
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	return nil
}
