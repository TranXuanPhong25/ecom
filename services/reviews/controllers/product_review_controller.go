package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/services/product-reviews/dtos"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/services"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/utils"
	"github.com/TranXuanPhong25/ecom/services/product-reviews/validators"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	service *services.ReviewService
}

func NewReviewController() *ReviewController {
	return &ReviewController{
		service: services.NewReviewService(),
	}
}

// CreateReview creates a new review
func (ctrl *ReviewController) CreateReview(c echo.Context) error {
	var req dtos.CreateReviewRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequestError(c, "Invalid request body")
	}

	// Validate request
	if err := validators.ValidateCreateReview(&req); err != nil {
		return utils.BadRequestError(c, validators.GetValidationErrors(err))
	}

	// Create review
	review, err := ctrl.service.CreateReview(&req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Review created successfully", review)
}

// GetReview gets a single review by ID
func (ctrl *ReviewController) GetReview(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return utils.BadRequestError(c, "Invalid review ID")
	}

	review, err := ctrl.service.GetReview(uint(id))
	if err != nil {
		return utils.NotFoundError(c, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Review retrieved successfully", review)
}

// GetProductReviews gets all reviews for a product with pagination
func (ctrl *ReviewController) GetProductReviews(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("productId"), 10, 32)
	if err != nil {
		return utils.BadRequestError(c, "Invalid product ID")
	}

	// Get pagination params
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	reviews, err := ctrl.service.GetProductReviews(uint(productID), page, limit)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Reviews retrieved successfully", reviews)
}

// GetUserReviews gets all reviews by a user
func (ctrl *ReviewController) GetUserReviews(c echo.Context) error {
	userID := c.Request().Header["X-User-Id"][0]
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	safeUserID, err := uuid.Parse(userID)

	if err != nil {
		return utils.BadRequestError(c, "Invalid user ID")
	}

	reviews, err := ctrl.service.GetUserReviews(safeUserID)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Reviews retrieved successfully", reviews)
}

// UpdateReview updates an existing review
func (ctrl *ReviewController) UpdateReview(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return utils.BadRequestError(c, "Invalid review ID")
	}

	var req dtos.UpdateReviewRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequestError(c, "Invalid request body")
	}

	// Validate request
	if err := validators.ValidateUpdateReview(&req); err != nil {
		return utils.BadRequestError(c, validators.GetValidationErrors(err))
	}

	// TODO: Extract userID from JWT token
	// For now, get from request body or header
	userIDStr := c.Request().Header.Get("X-User-Id")
	if userIDStr == "" {
		return utils.BadRequestError(c, "User ID is required")
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.BadRequestError(c, "Invalid User ID format")
	}

	review, err := ctrl.service.UpdateReview(uint(id), &req, userID)
	if err != nil {
		if err.Error() == "review not found" {
			return utils.NotFoundError(c, err.Error())
		}
		if err.Error() == "you can only update your own reviews" {
			return utils.ForbiddenError(c, err.Error())
		}
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Review updated successfully", review)
}

// DeleteReview deletes a review
func (ctrl *ReviewController) DeleteReview(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return utils.BadRequestError(c, "Invalid review ID")
	}

	// TODO: Extract userID from JWT token
	// For now, get from header
	userIDStr := c.Request().Header.Get("X-User-Id")
	if userIDStr == "" {
		return utils.BadRequestError(c, "User ID is required")
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.BadRequestError(c, "Invalid User ID format")
	}

	err = ctrl.service.DeleteReview(uint(id), userID)
	if err != nil {
		if err.Error() == "review not found" {
			return utils.NotFoundError(c, err.Error())
		}
		if err.Error() == "you can only delete your own reviews" {
			return utils.ForbiddenError(c, err.Error())
		}
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Review deleted successfully", nil)
}

// GetProductStats gets statistics for a product's reviews
func (ctrl *ReviewController) GetProductStats(c echo.Context) error {
	fmt.Printf("ổn mà")
	productID, err := strconv.ParseUint(c.Param("productId"), 10, 32)
	if err != nil {
		return utils.BadRequestError(c, "Invalid product ID")
	}

	stats, err := ctrl.service.GetProductStats(uint(productID))
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Statistics retrieved successfully", stats)
}
