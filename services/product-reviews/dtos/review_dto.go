package dtos

import "github.com/google/uuid"

// CreateReviewRequest represents the request body for creating a review
type CreateReviewRequest struct {
	ProductID uint      `json:"productId" validate:"required,min=1"`
	UserID    uuid.UUID `json:"userId" validate:"required"` // TODO: Extract from JWT token
	Username  string    `json:"username" validate:"required,min=2,max=50"`
	Rating    int       `json:"rating" validate:"required,min=1,max=5"`
	Title     string    `json:"title" validate:"required,min=10,max=100"`
	Comment   string    `json:"comment" validate:"required,min=20,max=1000"`
}

// UpdateReviewRequest represents the request body for updating a review
type UpdateReviewRequest struct {
	Rating  int    `json:"rating" validate:"required,min=1,max=5"`
	Title   string `json:"title" validate:"required,min=10,max=100"`
	Comment string `json:"comment" validate:"required,min=20,max=1000"`
}

// ReviewResponse represents a review in API responses
type ReviewResponse struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"productId"`
	UserID    uuid.UUID `json:"userId"`
	Username  string    `json:"username"`
	Rating    int       `json:"rating"`
	Title     string    `json:"title"`
	Comment   string    `json:"comment"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
}

// ReviewListResponse represents a paginated list of reviews
type ReviewListResponse struct {
	Reviews []ReviewResponse `json:"reviews"`
	Page    int              `json:"page"`
	Limit   int              `json:"limit"`
	Total   int64            `json:"total"` // TODO: Implement real count query
}

// ProductRatingStatsResponse represents product rating statistics
type ProductRatingStatsResponse struct {
	ProductID      uint    `json:"productId"`
	AverageRating  float64 `json:"averageRating"`
	TotalReviews   int64   `json:"totalReviews"`
	FiveStarCount  int64   `json:"fiveStarCount"`
	FourStarCount  int64   `json:"fourStarCount"`
	ThreeStarCount int64   `json:"threeStarCount"`
	TwoStarCount   int64   `json:"twoStarCount"`
	OneStarCount   int64   `json:"oneStarCount"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
