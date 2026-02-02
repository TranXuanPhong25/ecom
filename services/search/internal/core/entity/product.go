package entity

import "time"

// Product represents a product in the domain model
// This is a pure domain entity with no external dependencies
type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"categoryId"`
	CoverImage  string    `json:"coverImage"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// SearchResult represents the result of a search operation
type SearchResult struct {
	Products []Product
	Total    int64
	Page     int
	Limit    int
}
