package dtos

// CreateReviewRequest represents the request body for creating a review
type CreateReviewRequest struct {
	ProductID uint   `json:"product_id" validate:"required,min=1"`
	UserID    uint   `json:"user_id" validate:"required,min=1"` // TODO: Extract from JWT token
	Username  string `json:"username" validate:"required,min=2,max=50"`
	Rating    int    `json:"rating" validate:"required,min=1,max=5"`
	Title     string `json:"title" validate:"required,min=10,max=100"`
	Comment   string `json:"comment" validate:"required,min=20,max=1000"`
}

// UpdateReviewRequest represents the request body for updating a review
type UpdateReviewRequest struct {
	Rating  int    `json:"rating" validate:"required,min=1,max=5"`
	Title   string `json:"title" validate:"required,min=10,max=100"`
	Comment string `json:"comment" validate:"required,min=20,max=1000"`
}

// ReviewResponse represents a review in API responses
type ReviewResponse struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Rating    int    `json:"rating"`
	Title     string `json:"title"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
	ProductID      uint    `json:"product_id"`
	AverageRating  float64 `json:"average_rating"`
	TotalReviews   int64   `json:"total_reviews"`
	FiveStarCount  int64   `json:"five_star_count"`
	FourStarCount  int64   `json:"four_star_count"`
	ThreeStarCount int64   `json:"three_star_count"`
	TwoStarCount   int64   `json:"two_star_count"`
	OneStarCount   int64   `json:"one_star_count"`
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
