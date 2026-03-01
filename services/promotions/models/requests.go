package models

import (
	"time"

	"github.com/google/uuid"
)

// CreateEventBannerRequest - Request để tạo event banner
type CreateEventBannerRequest struct {
	Title       string    `json:"title" validate:"required,min=3,max=200"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageUrl" validate:"required,url"`
	LinkURL     string    `json:"linkUrl" validate:"omitempty,url"`
	StartTime   time.Time `json:"startTime" validate:"required"`
	EndTime     time.Time `json:"endTime" validate:"required,gtfield=StartTime"`
	Priority    int       `json:"priority"`
	IsActive    bool      `json:"isActive"`
	EventType   string    `json:"eventType" validate:"required,oneof=black_friday flash_sale new_year holiday seasonal other"`
	Position    string    `json:"position" validate:"omitempty,oneof=main sidebar popup"`
}

// UpdateEventBannerRequest - Request để cập nhật event banner
type UpdateEventBannerRequest struct {
	ID          uuid.UUID  `json:"id" validate:"required,uuid"`
	Title       string     `json:"title" validate:"omitempty,min=3,max=200"`
	Description string     `json:"description"`
	ImageURL    string     `json:"imageUrl" validate:"omitempty,url"`
	LinkURL     string     `json:"linkUrl" validate:"omitempty,url"`
	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	Priority    *int       `json:"priority"`
	IsActive    *bool      `json:"isActive"`
	EventType   string     `json:"eventType" validate:"omitempty,oneof=black_friday flash_sale new_year holiday seasonal other"`
	Position    string     `json:"position" validate:"omitempty,oneof=main sidebar popup"`
}

// CreatePromoBarRequest - Request để tạo promo bar
type CreatePromoBarRequest struct {
	Message         string    `json:"message" validate:"required,min=3,max=300"`
	BackgroundColor string    `json:"backgroundColor" validate:"omitempty,hexcolor"`
	TextColor       string    `json:"textColor" validate:"omitempty,hexcolor"`
	LinkURL         string    `json:"linkUrl" validate:"omitempty,url"`
	StartTime       time.Time `json:"startTime" validate:"required"`
	EndTime         time.Time `json:"endTime" validate:"required,gtfield=StartTime"`
	IsActive        bool      `json:"isActive"`
	Priority        int       `json:"priority"`
	IsCloseable     bool      `json:"isCloseable"`
}

// UpdatePromoBarRequest - Request để cập nhật promo bar
type UpdatePromoBarRequest struct {
	ID              uuid.UUID  `json:"id" validate:"required,uuid"`
	Message         string     `json:"message" validate:"omitempty,min=3,max=300"`
	BackgroundColor string     `json:"backgroundColor" validate:"omitempty,hexcolor"`
	TextColor       string     `json:"textColor" validate:"omitempty,hexcolor"`
	LinkURL         string     `json:"linkUrl" validate:"omitempty,url"`
	StartTime       *time.Time `json:"startTime"`
	EndTime         *time.Time `json:"endTime"`
	IsActive        *bool      `json:"isActive"`
	Priority        *int       `json:"priority"`
	IsCloseable     *bool      `json:"isCloseable"`
}
