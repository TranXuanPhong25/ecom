package dto

import (
	"time"

	"github.com/rengumin/fulfillment/internal/core/entity"
)

// Request to schedule a pickup
type SchedulePickupRequest struct {
	OrderID              int64          `json:"orderId" binding:"required"`
	ShopID               string         `json:"shopId" binding:"required"`
	PickupAddress        string         `json:"pickupAddress" binding:"required"`
	PickupContactName    string         `json:"pickupContactName"`
	PickupContactPhone   string         `json:"pickupContactPhone"`
	DeliveryAddress      string         `json:"deliveryAddress" binding:"required"`
	DeliveryContactName  string         `json:"deliveryContactName" binding:"required"`
	DeliveryContactPhone string         `json:"deliveryContactPhone" binding:"required"`
	WeightGrams          int            `json:"weightGrams"`
	Dimensions           map[string]any `json:"dimensions"`
	SpecialInstructions  string         `json:"specialInstructions"`
}

type SchedulePickupResponse struct {
	PackageNumber     string    `json:"packageNumber"`
	PickupScheduledAt time.Time `json:"pickupScheduledAt"`
	EstimatedDelivery time.Time `json:"estimatedDelivery"`
	Message           string    `json:"message"`
}

// Update package location
type UpdateLocationRequest struct {
	PackageNumber string `json:"packageNumber" binding:"required"`
	Location      string `json:"location" binding:"required"`
	ScannedAt     string `json:"scannedAt"`
}

// Mark package as picked up
type MarkPickedUpRequest struct {
	PackageNumber string `json:"packageNumber" binding:"required"`
	PickupBy      string `json:"pickupBy"`
	Notes         string `json:"notes"`
}

// Update delivery status
type UpdateDeliveryStatusRequest struct {
	PackageNumber        string              `json:"packageNumber" binding:"required"`
	Status               entity.PackageStatus `json:"status" binding:"required"` // DELIVERED, DELIVERY_FAILED
	DeliverySignatureURL string              `json:"deliverySignatureUrl"`
	FailureReason        string              `json:"failureReason"`
	AttemptedAt          string              `json:"attemptedAt"`
}

// Package tracking response
type PackageTrackingDTO struct {
	PackageNumber     string             `json:"packageNumber"`
	OrderID           int64              `json:"orderId"`
	Status            string             `json:"status"`
	CurrentLocation   string             `json:"currentLocation,omitempty"`
	LastScanAt        *time.Time         `json:"lastScanAt,omitempty"`
	EstimatedDelivery *time.Time         `json:"estimatedDelivery,omitempty"`
	DeliveryAttempts  int                `json:"deliveryAttempts"`
	TrackingEvents    []TrackingEventDTO `json:"trackingEvents"`
	CreatedAt         time.Time          `json:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt"`
}

type TrackingEventDTO struct {
	Location  string    `json:"location"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Notes     string    `json:"notes,omitempty"`
}

// List packages query params
type ListPackagesQuery struct {
	ShopID   string              `form:"shopId"`
	Status   entity.PackageStatus `form:"status"`
	Zone     string              `form:"zone"`
	Page     int                 `form:"page" binding:"min=1"`
	PageSize int                 `form:"pageSize" binding:"min=1,max=100"`
}

type PackageListDTO struct {
	ID                int64      `json:"id"`
	PackageNumber     string     `json:"packageNumber"`
	OrderID           int64      `json:"orderId"`
	ShopID            string     `json:"shopId"`
	Status            string     `json:"status"`
	PickupScheduledAt *time.Time `json:"pickupScheduledAt"`
	EstimatedDelivery *time.Time `json:"estimatedDelivery"`
	DeliveryZone      string     `json:"deliveryZone,omitempty"`
	CreatedAt         time.Time  `json:"createdAt"`
}

type PageResponse struct {
	Content       []PackageListDTO `json:"content"`
	TotalElements int64            `json:"totalElements"`
	TotalPages    int              `json:"totalPages"`
	PageNumber    int              `json:"pageNumber"`
	PageSize      int              `json:"pageSize"`
}
