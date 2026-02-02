package handler

import (
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/port/in"
	"github.com/labstack/echo/v4"
)

// SearchHandler handles HTTP requests for search operations
// This is the input adapter for HTTP
type SearchHandler struct {
	searchService in.SearchServicePort
}

// NewSearchHandler creates a new SearchHandler with dependency injection
func NewSearchHandler(searchService in.SearchServicePort) *SearchHandler {
	return &SearchHandler{
		searchService: searchService,
	}
}

// SearchProducts handles product search requests
// GET /search?q=keyword&page=1&limit=20
func (h *SearchHandler) SearchProducts(c echo.Context) error {
	// Parse query parameters
	query := c.QueryParam("q")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Create request DTO
	request := dto.SearchRequest{
		Query: query,
		Page:  page,
		Limit: limit,
	}

	// Validate request
	if request.Query == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Bad Request",
			Message: "Search query is required",
		})
	}

	// Call service
	response, err := h.searchService.SearchProducts(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   "Internal Server Error",
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Products found successfully",
		Data:    response,
	})
}
