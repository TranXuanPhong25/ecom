package dtos

import (
	"time"

	"github.com/google/uuid"
)

// EventBannerDTO - DTO cho Event Banner
type EventBannerDTO struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	ImageURL    string    `json:"imageUrl"`
	LinkURL     string    `json:"linkUrl,omitempty"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Priority    int       `json:"priority"`
	IsActive    bool      `json:"isActive"`
	EventType   string    `json:"eventType"`
	Position    string    `json:"position"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// PromoBarDTO - DTO cho Promo Bar
type PromoBarDTO struct {
	ID              uuid.UUID `json:"id"`
	Message         string    `json:"message"`
	BackgroundColor string    `json:"backgroundColor"`
	TextColor       string    `json:"textColor"`
	LinkURL         string    `json:"linkUrl,omitempty"`
	StartTime       time.Time `json:"startTime"`
	EndTime         time.Time `json:"endTime"`
	IsActive        bool      `json:"isActive"`
	Priority        int       `json:"priority"`
	IsCloseable     bool      `json:"isCloseable"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// GetActivePromotionsResponse - Response cho API lấy promotions đang active
type GetActivePromotionsResponse struct {
	EventBanners []EventBannerDTO `json:"eventBanners"`
	PromoBars    []PromoBarDTO    `json:"promoBars"`
}
