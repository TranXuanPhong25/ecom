package services

import (
	"errors"
	"time"

	"github.com/TranXuanPhong25/ecom/services/product-reviews/dtos"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/models"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/repositories"
	"github.com/google/uuid"
)

type ReviewService struct {
	repo *repositories.ReviewRepository
}

func NewReviewService() *ReviewService {
	return &ReviewService{
		repo: repositories.NewReviewRepository(),
	}
}

// CreateReview creates a new review
func (s *ReviewService) CreateReview(req *dtos.CreateReviewRequest) (*dtos.ReviewResponse, error) {
	review := &models.Review{
		ProductID: req.ProductID,
		UserID:    req.UserID,
		Username:  req.Username,
		Rating:    req.Rating,
		Title:     req.Title,
		Comment:   req.Comment,
	}

	err := s.repo.Create(review)
	if err != nil {
		return nil, err
	}

	return s.modelToResponse(review), nil
}

// GetReview gets a review by ID
func (s *ReviewService) GetReview(id uint) (*dtos.ReviewResponse, error) {
	review, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("review not found")
	}
	return s.modelToResponse(review), nil
}

// GetProductReviews gets reviews for a product with pagination
func (s *ReviewService) GetProductReviews(productID uint, page, limit int) (*dtos.ReviewListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20 // default
	}

	offset := (page - 1) * limit
	reviews, err := s.repo.FindByProductID(productID, limit, offset)
	if err != nil {
		return nil, err
	}

	// TODO: Implement real count query
	total, _ := s.repo.CountByProductID(productID)

	reviewResponses := make([]dtos.ReviewResponse, len(reviews))
	for i, review := range reviews {
		reviewResponses[i] = *s.modelToResponse(&review)
	}

	return &dtos.ReviewListResponse{
		Reviews: reviewResponses,
		Page:    page,
		Limit:   limit,
		Total:   total, // -1 for now
	}, nil
}

// GetUserReviews gets all reviews by a user
func (s *ReviewService) GetUserReviews(userID uuid.UUID) ([]dtos.ReviewResponse, error) {
	reviews, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	reviewResponses := make([]dtos.ReviewResponse, len(reviews))
	for i, review := range reviews {
		reviewResponses[i] = *s.modelToResponse(&review)
	}

	return reviewResponses, nil
}

// UpdateReview updates an existing review
func (s *ReviewService) UpdateReview(id uint, req *dtos.UpdateReviewRequest, userID uuid.UUID) (*dtos.ReviewResponse, error) {
	review, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("review not found")
	}

	// Check ownership
	if review.UserID != userID {
		return nil, errors.New("you can only update your own reviews")
	}

	// Update fields
	review.Rating = req.Rating
	review.Title = req.Title
	review.Comment = req.Comment

	err = s.repo.Update(review)
	if err != nil {
		return nil, err
	}

	return s.modelToResponse(review), nil
}

// DeleteReview deletes a review
func (s *ReviewService) DeleteReview(id uint, userID uuid.UUID) error {
	review, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("review not found")
	}

	// Check ownership
	if review.UserID != userID {
		return errors.New("you can only delete your own reviews")
	}

	return s.repo.Delete(id)
}

func (s *ReviewService) GetProductStats(productID uint) (*dtos.ProductRatingStatsResponse, error) {
	return s.repo.GetProductStats(productID)
}

// modelToResponse converts a Review model to ReviewResponse DTO
func (s *ReviewService) modelToResponse(review *models.Review) *dtos.ReviewResponse {
	return &dtos.ReviewResponse{
		ID:        review.ID,
		ProductID: review.ProductID,
		UserID:    review.UserID,
		Username:  review.Username,
		Rating:    review.Rating,
		Title:     review.Title,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt.Format(time.RFC3339),
		UpdatedAt: review.UpdatedAt.Format(time.RFC3339),
	}
}
