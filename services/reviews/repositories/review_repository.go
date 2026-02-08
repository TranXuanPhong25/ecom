package repositories

import (
	"github.com/TranXuanPhong25/ecom/services/product-reviews/database"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/dtos"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{
		db: database.DB,
	}
}

// Create creates a new review
func (r *ReviewRepository) Create(review *models.Review) error {
	return r.db.Create(review).Error
}

// FindByID finds a review by ID
func (r *ReviewRepository) FindByID(id uint) (*models.Review, error) {
	var review models.Review
	err := r.db.First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// FindByProductID finds reviews by product ID with pagination
func (r *ReviewRepository) FindByProductID(productID uint, limit, offset int) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("product_id = ?", productID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&reviews).Error
	return reviews, err
}

// FindByUserID finds all reviews by user ID
func (r *ReviewRepository) FindByUserID(userID uuid.UUID) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&reviews).Error
	return reviews, err
}

// Update updates an existing review
func (r *ReviewRepository) Update(review *models.Review) error {
	return r.db.Save(review).Error
}

// Delete deletes a review by ID
func (r *ReviewRepository) Delete(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}

// CountByProductID counts reviews for a product (TODO: implement for real pagination)
func (r *ReviewRepository) CountByProductID(productID uint) (int64, error) {
	// TODO: Implement real count query
	// For now, return a placeholder value
	var count int64
	r.db.Where("product_id = ?", productID).Count(&count)
	return count, nil
}
func (r *ReviewRepository) GetProductStats(productID uint) (*dtos.ProductRatingStatsResponse, error) {
	var stats dtos.ProductRatingStatsResponse

	err := r.db.Model(&models.Review{}).
		Select(`
			SUM(CASE WHEN rating = 5 THEN 1 ELSE 0 END) AS five_star_count,
			SUM(CASE WHEN rating = 4 THEN 1 ELSE 0 END) AS four_star_count,
			SUM(CASE WHEN rating = 3 THEN 1 ELSE 0 END) AS three_star_count,
			SUM(CASE WHEN rating = 2 THEN 1 ELSE 0 END) AS two_star_count,
			SUM(CASE WHEN rating = 1 THEN 1 ELSE 0 END) AS one_star_count,
			COUNT(*) AS total_reviews,
			COALESCE(AVG(rating), 0) AS average_rating
		`).
		Where("product_id = ?", productID).
		Scan(&stats).Error

	if err != nil {
		return nil, err
	}

	stats.ProductID = productID
	return &stats, nil
}
