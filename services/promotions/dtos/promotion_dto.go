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
	ImageURL    string    `json:"image_url"`
	LinkURL     string    `json:"link_url,omitempty"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Priority    int       `json:"priority"`
	IsActive    bool      `json:"is_active"`
	EventType   string    `json:"event_type"`
	Position    string    `json:"position"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PromoBarDTO - DTO cho Promo Bar
type PromoBarDTO struct {
	ID              uuid.UUID `json:"id"`
	Message         string    `json:"message"`
	BackgroundColor string    `json:"background_color"`
	TextColor       string    `json:"text_color"`
	LinkURL         string    `json:"link_url,omitempty"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	IsActive        bool      `json:"is_active"`
	Priority        int       `json:"priority"`
	IsCloseable     bool      `json:"is_closeable"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GetActivePromotionsResponse - Response cho API lấy promotions đang active
type GetActivePromotionsResponse struct {
	EventBanners []EventBannerDTO `json:"event_banners"`
	PromoBars    []PromoBarDTO    `json:"promo_bars"`
}
