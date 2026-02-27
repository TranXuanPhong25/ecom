package port

import (
	"context"
	"github.com/rengumin/fulfillment/internal/core/dto"
	"github.com/rengumin/fulfillment/internal/core/entity"
)

// FulfillmentService defines business logic for fulfillment operations
type FulfillmentService interface {
	// Schedule pickup from seller
	SchedulePickup(ctx context.Context, req dto.SchedulePickupRequest) (*dto.SchedulePickupResponse, error)
	
	// Mark package as picked up by driver
	MarkPickedUp(ctx context.Context, req dto.MarkPickedUpRequest) error
	
	// Update package location during transit
	UpdateLocation(ctx context.Context, req dto.UpdateLocationRequest) error
	
	// Update delivery status (delivered or failed)
	UpdateDeliveryStatus(ctx context.Context, req dto.UpdateDeliveryStatusRequest) error
	
	// Get package tracking info
	GetPackageTracking(ctx context.Context, packageNumber string) (*dto.PackageTrackingDTO, error)
	
	// List packages with filters
	ListPackages(ctx context.Context, query dto.ListPackagesQuery) (*dto.PageResponse, error)
	
	// Get package by order ID
	GetPackageByOrderID(ctx context.Context, orderID int64) (*entity.FulfillmentPackage, error)
}

// EventPublisher publishes events to Kafka
type EventPublisher interface {
	PublishPickupScheduled(ctx context.Context, pkg *entity.FulfillmentPackage) error
	PublishPickedUp(ctx context.Context, pkg *entity.FulfillmentPackage) error
	PublishInTransit(ctx context.Context, pkg *entity.FulfillmentPackage) error
	PublishOutForDelivery(ctx context.Context, pkg *entity.FulfillmentPackage) error
	PublishDelivered(ctx context.Context, pkg *entity.FulfillmentPackage) error
	PublishDeliveryFailed(ctx context.Context, pkg *entity.FulfillmentPackage) error
}
