package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rengumin/fulfillment/internal/config"
	"github.com/rengumin/fulfillment/internal/core/dto"
	"github.com/rengumin/fulfillment/internal/core/entity"
	"github.com/rengumin/fulfillment/internal/core/port"
)

type fulfillmentService struct {
	repo      port.PackageRepository
	publisher port.EventPublisher
	config    *config.Config
}

func NewFulfillmentService(repo port.PackageRepository, publisher port.EventPublisher, cfg *config.Config) port.FulfillmentService {
	return &fulfillmentService{
		repo:      repo,
		publisher: publisher,
		config:    cfg,
	}
}

// SchedulePickup schedules a pickup from seller
func (s *fulfillmentService) SchedulePickup(ctx context.Context, req dto.SchedulePickupRequest) (*dto.SchedulePickupResponse, error) {
	// Generate unique package number
	packageNumber := generatePackageNumber()

	// Calculate pickup schedule (next day, 9-11 AM by default)
	pickupTime := time.Now().Add(time.Duration(s.config.DefaultPickupWindow) * time.Hour)

	// Calculate estimated delivery (3 days from pickup by default)
	estimatedDelivery := pickupTime.Add(time.Duration(s.config.EstimatedDeliveryDays) * 24 * time.Hour)

	// Determine delivery zone from address (simplified logic)
	deliveryZone := determineDeliveryZone(req.DeliveryAddress)

	// Create fulfillment package
	pkg := &entity.FulfillmentPackage{
		PackageNumber:        packageNumber,
		OrderID:              req.OrderID,
		ShopID:               req.ShopID,
		PickupAddress:        req.PickupAddress,
		PickupContactName:    strPtr(req.PickupContactName),
		PickupContactPhone:   strPtr(req.PickupContactPhone),
		PickupScheduledAt:    &pickupTime,
		Status:               entity.StatusPendingPickup,
		DeliveryAddress:      req.DeliveryAddress,
		DeliveryContactName:  req.DeliveryContactName,
		DeliveryContactPhone: req.DeliveryContactPhone,
		DeliveryZone:         &deliveryZone,
		EstimatedDelivery:    &estimatedDelivery,
		WeightGrams:          intPtr(req.WeightGrams),
		Dimensions:           &req.Dimensions,
		SpecialInstructions:  strPtr(req.SpecialInstructions),
	}

	// Save to database
	if err := s.repo.Create(ctx, pkg); err != nil {
		return nil, fmt.Errorf("failed to create package: %w", err)
	}

	// Publish pickup scheduled event
	if err := s.publisher.PublishPickupScheduled(ctx, pkg); err != nil {
		// Log error but don't fail the operation
		fmt.Printf("failed to publish pickup scheduled event: %v\n", err)
	}

	return &dto.SchedulePickupResponse{
		PackageNumber:     packageNumber,
		PickupScheduledAt: pickupTime,
		EstimatedDelivery: estimatedDelivery,
		Message:           "Pickup scheduled successfully. Driver will arrive at scheduled time.",
	}, nil
}

// MarkPickedUp marks package as picked up by driver
func (s *fulfillmentService) MarkPickedUp(ctx context.Context, req dto.MarkPickedUpRequest) error {
	pkg, err := s.repo.FindByPackageNumber(ctx, req.PackageNumber)
	if err != nil {
		return fmt.Errorf("package not found: %w", err)
	}

	// Validate current status
	if pkg.Status != entity.StatusPendingPickup {
		return fmt.Errorf("invalid status transition: current status is %s", pkg.Status)
	}

	// Update package status
	now := time.Now()
	pkg.Status = entity.StatusPickedUp
	pkg.PickupCompletedAt = &now
	pkg.CurrentHubLocation = strPtr("Pickup Completed")
	pkg.LastScanAt = &now
	pkg.UpdatedAt = now

	if err := s.repo.Update(ctx, pkg); err != nil {
		return fmt.Errorf("failed to update package: %w", err)
	}

	// Publish picked up event
	if err := s.publisher.PublishPickedUp(ctx, pkg); err != nil {
		fmt.Printf("failed to publish picked up event: %v\n", err)
	}

	return nil
}

// UpdateLocation updates package location during transit
func (s *fulfillmentService) UpdateLocation(ctx context.Context, req dto.UpdateLocationRequest) error {
	pkg, err := s.repo.FindByPackageNumber(ctx, req.PackageNumber)
	if err != nil {
		return fmt.Errorf("package not found: %w", err)
	}

	// Parse scanned time
	var scannedAt time.Time
	if req.ScannedAt != "" {
		scannedAt, _ = time.Parse(time.RFC3339, req.ScannedAt)
	} else {
		scannedAt = time.Now()
	}

	// Update location
	pkg.CurrentHubLocation = &req.Location
	pkg.LastScanAt = &scannedAt
	pkg.UpdatedAt = time.Now()

	// Update status based on location
	if pkg.Status == entity.StatusPickedUp {
		pkg.Status = entity.StatusAtHub
	} else if pkg.Status == entity.StatusAtHub {
		pkg.Status = entity.StatusInTransit
	}

	if err := s.repo.Update(ctx, pkg); err != nil {
		return fmt.Errorf("failed to update package: %w", err)
	}

	// Publish in transit event
	if pkg.Status == entity.StatusInTransit {
		if err := s.publisher.PublishInTransit(ctx, pkg); err != nil {
			fmt.Printf("failed to publish in transit event: %v\n", err)
		}
	}

	return nil
}

// UpdateDeliveryStatus updates delivery status (delivered or failed)
func (s *fulfillmentService) UpdateDeliveryStatus(ctx context.Context, req dto.UpdateDeliveryStatusRequest) error {
	pkg, err := s.repo.FindByPackageNumber(ctx, req.PackageNumber)
	if err != nil {
		return fmt.Errorf("package not found: %w", err)
	}

	now := time.Now()

	if req.Status == entity.StatusDelivered {
		pkg.Status = entity.StatusDelivered
		pkg.DeliveredAt = &now
		pkg.DeliverySignatureURL = strPtr(req.DeliverySignatureURL)

		if err := s.repo.Update(ctx, pkg); err != nil {
			return fmt.Errorf("failed to update package: %w", err)
		}

		// Publish delivered event
		if err := s.publisher.PublishDelivered(ctx, pkg); err != nil {
			fmt.Printf("failed to publish delivered event: %v\n", err)
		}

	} else if req.Status == entity.StatusDeliveryFailed {
		pkg.DeliveryAttempts++
		pkg.LastDeliveryAttemptAt = &now
		pkg.DeliveryFailureReason = strPtr(req.FailureReason)

		// Check if max attempts reached
		if pkg.DeliveryAttempts >= s.config.MaxDeliveryAttempts {
			pkg.Status = entity.StatusReturnedToSeller
		} else {
			pkg.Status = entity.StatusDeliveryFailed
		}

		if err := s.repo.Update(ctx, pkg); err != nil {
			return fmt.Errorf("failed to update package: %w", err)
		}

		// Publish delivery failed event
		if err := s.publisher.PublishDeliveryFailed(ctx, pkg); err != nil {
			fmt.Printf("failed to publish delivery failed event: %v\n", err)
		}
	}

	return nil
}

// GetPackageTracking gets package tracking info
func (s *fulfillmentService) GetPackageTracking(ctx context.Context, packageNumber string) (*dto.PackageTrackingDTO, error) {
	pkg, err := s.repo.FindByPackageNumber(ctx, packageNumber)
	if err != nil {
		return nil, fmt.Errorf("package not found: %w", err)
	}

	// Build tracking events (simplified - in production, would have separate tracking_events table)
	events := buildTrackingEvents(pkg)

	tracking := &dto.PackageTrackingDTO{
		PackageNumber:     pkg.PackageNumber,
		OrderID:           pkg.OrderID,
		Status:            string(pkg.Status),
		CurrentLocation:   safeString(pkg.CurrentHubLocation),
		LastScanAt:        pkg.LastScanAt,
		EstimatedDelivery: pkg.EstimatedDelivery,
		DeliveryAttempts:  pkg.DeliveryAttempts,
		TrackingEvents:    events,
		CreatedAt:         pkg.CreatedAt,
		UpdatedAt:         pkg.UpdatedAt,
	}

	return tracking, nil
}

// ListPackages lists packages with filters
func (s *fulfillmentService) ListPackages(ctx context.Context, query dto.ListPackagesQuery) (*dto.PageResponse, error) {
	packages, total, err := s.repo.FindAll(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list packages: %w", err)
	}

	// Convert to DTOs
	content := make([]dto.PackageListDTO, len(packages))
	for i, pkg := range packages {
		content[i] = dto.PackageListDTO{
			ID:                pkg.ID,
			PackageNumber:     pkg.PackageNumber,
			OrderID:           pkg.OrderID,
			ShopID:            pkg.ShopID,
			Status:            string(pkg.Status),
			PickupScheduledAt: pkg.PickupScheduledAt,
			EstimatedDelivery: pkg.EstimatedDelivery,
			DeliveryZone:      safeString(pkg.DeliveryZone),
			CreatedAt:         pkg.CreatedAt,
		}
	}

	totalPages := int(total) / query.PageSize
	if int(total)%query.PageSize > 0 {
		totalPages++
	}

	return &dto.PageResponse{
		Content:       content,
		TotalElements: total,
		TotalPages:    totalPages,
		PageNumber:    query.Page,
		PageSize:      query.PageSize,
	}, nil
}

// GetPackageByOrderID gets package by order ID
func (s *fulfillmentService) GetPackageByOrderID(ctx context.Context, orderID int64) (*entity.FulfillmentPackage, error) {
	return s.repo.FindByOrderID(ctx, orderID)
}

// Helper functions

func generatePackageNumber() string {
	return fmt.Sprintf("PKG%d%06d", time.Now().Unix(), rand.Intn(1000000))
}

func determineDeliveryZone(address string) string {
	// Simplified zone determination logic
	// In production, would use geocoding or address parsing
	if len(address) > 0 {
		if address[0] >= 'A' && address[0] <= 'M' {
			return "ZONE_NORTH"
		}
		return "ZONE_SOUTH"
	}
	return "ZONE_UNKNOWN"
}

func buildTrackingEvents(pkg *entity.FulfillmentPackage) []dto.TrackingEventDTO {
	events := []dto.TrackingEventDTO{}

	if pkg.PickupScheduledAt != nil {
		events = append(events, dto.TrackingEventDTO{
			Location:  "Pickup Scheduled",
			Status:    string(entity.StatusPendingPickup),
			Timestamp: *pkg.PickupScheduledAt,
		})
	}

	if pkg.PickupCompletedAt != nil {
		events = append(events, dto.TrackingEventDTO{
			Location:  "Picked Up from Seller",
			Status:    string(entity.StatusPickedUp),
			Timestamp: *pkg.PickupCompletedAt,
		})
	}

	if pkg.LastScanAt != nil && pkg.CurrentHubLocation != nil {
		events = append(events, dto.TrackingEventDTO{
			Location:  *pkg.CurrentHubLocation,
			Status:    string(pkg.Status),
			Timestamp: *pkg.LastScanAt,
		})
	}

	if pkg.DeliveredAt != nil {
		events = append(events, dto.TrackingEventDTO{
			Location:  "Delivered to Customer",
			Status:    string(entity.StatusDelivered),
			Timestamp: *pkg.DeliveredAt,
		})
	}

	return events
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func intPtr(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
