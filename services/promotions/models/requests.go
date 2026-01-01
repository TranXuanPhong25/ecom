package models

import (
	"time"

	"github.com/google/uuid"
)

// CreateEventBannerRequest - Request để tạo event banner
type CreateEventBannerRequest struct {
	Title       string    `json:"title" validate:"required,min=3,max=200"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url" validate:"required,url"`
	LinkURL     string    `json:"link_url" validate:"omitempty,url"`
	StartTime   time.Time `json:"start_time" validate:"required"`
	EndTime     time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	Priority    int       `json:"priority"`
	IsActive    bool      `json:"is_active"`
	EventType   string    `json:"event_type" validate:"required,oneof=black_friday flash_sale new_year holiday seasonal other"`
	Position    string    `json:"position" validate:"omitempty,oneof=main sidebar popup"`
}

// UpdateEventBannerRequest - Request để cập nhật event banner
type UpdateEventBannerRequest struct {
	ID          uuid.UUID  `json:"id" validate:"required,uuid"`
	Title       string     `json:"title" validate:"omitempty,min=3,max=200"`
	Description string     `json:"description"`
	ImageURL    string     `json:"image_url" validate:"omitempty,url"`
	LinkURL     string     `json:"link_url" validate:"omitempty,url"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	Priority    *int       `json:"priority"`
	IsActive    *bool      `json:"is_active"`
	EventType   string     `json:"event_type" validate:"omitempty,oneof=black_friday flash_sale new_year holiday seasonal other"`
	Position    string     `json:"position" validate:"omitempty,oneof=main sidebar popup"`
}

// CreatePromoBarRequest - Request để tạo promo bar
type CreatePromoBarRequest struct {
	Message         string    `json:"message" validate:"required,min=3,max=300"`
	BackgroundColor string    `json:"background_color" validate:"omitempty,hexcolor"`
	TextColor       string    `json:"text_color" validate:"omitempty,hexcolor"`
	LinkURL         string    `json:"link_url" validate:"omitempty,url"`
	StartTime       time.Time `json:"start_time" validate:"required"`
	EndTime         time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	IsActive        bool      `json:"is_active"`
	Priority        int       `json:"priority"`
	IsCloseable     bool      `json:"is_closeable"`
}

// UpdatePromoBarRequest - Request để cập nhật promo bar
type UpdatePromoBarRequest struct {
	ID              uuid.UUID  `json:"id" validate:"required,uuid"`
	Message         string     `json:"message" validate:"omitempty,min=3,max=300"`
	BackgroundColor string     `json:"background_color" validate:"omitempty,hexcolor"`
	TextColor       string     `json:"text_color" validate:"omitempty,hexcolor"`
	LinkURL         string     `json:"link_url" validate:"omitempty,url"`
	StartTime       *time.Time `json:"start_time"`
	EndTime         *time.Time `json:"end_time"`
	IsActive        *bool      `json:"is_active"`
	Priority        *int       `json:"priority"`
	IsCloseable     *bool      `json:"is_closeable"`
}
