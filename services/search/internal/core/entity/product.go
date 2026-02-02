package entity

import "time"

// Product represents a product in the domain model
// This is a pure domain entity with no external dependencies
type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  uint      `json:"category_id"`
	CoverImage  string    `json:"cover_image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SearchResult represents the result of a search operation
type SearchResult struct {
	Products []Product
	Total    int64
	Page     int
	Limit    int
}
