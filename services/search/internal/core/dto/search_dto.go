package dto

// SearchRequest represents a search query request
type SearchRequest struct {
	Query string `json:"query" validate:"required,min=1,max=200"`
	Page  int    `json:"page" validate:"min=1"`
	Limit int    `json:"limit" validate:"min=1,max=100"`
}

// SearchResponse represents the response of a search operation
type SearchResponse struct {
	Products []ProductDTO `json:"products"`
	Total    int64        `json:"total"`
	Page     int          `json:"page"`
	Limit    int          `json:"limit"`
}

// ProductDTO represents a product in API responses
type ProductDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"categoryId"`
	CoverImage  string  `json:"coverImage"`
}

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
