package service

import (
	"context"
	"log"
	"time"

	"github.com/rengumin/fulfillment/internal/core/dto"
	"github.com/rengumin/fulfillment/internal/core/entity"
	"github.com/rengumin/fulfillment/internal/core/port"
)

// AutoDeliverySimulator simulates delivery process automatically
type AutoDeliverySimulator struct {
	repo      port.PackageRepository
	publisher port.EventPublisher
}

func NewAutoDeliverySimulator(repo port.PackageRepository, publisher port.EventPublisher) *AutoDeliverySimulator {
	return &AutoDeliverySimulator{
		repo:      repo,
		publisher: publisher,
	}
}

// Run cron job - chạy mỗi 30 phút hoặc 1 giờ
func (s *AutoDeliverySimulator) Run(ctx context.Context) error {

	// Tìm tất cả packages đang trong quá trình giao hàng
	statuses := []entity.PackageStatus{
		entity.StatusPendingPickup,
		entity.StatusPickedUp,
		entity.StatusAtHub,
		entity.StatusInTransit,
		entity.StatusOutForDelivery,
	}

	for _, status := range statuses {
		packages, _, err := s.repo.FindAll(ctx, dto.ListPackagesQuery{
			Status:   status,
			PageSize: 100,
		})

		if err != nil {
			log.Printf("Error finding packages with status %s: %v", status, err)
			continue
		}

		for _, pkg := range packages {
			if err := s.progressPackage(ctx, &pkg); err != nil {
				log.Printf("Error progressing package %s: %v", pkg.PackageNumber, err)
			}
		}
	}

	return nil
}

// progressPackage đẩy package qua stage tiếp theo
func (s *AutoDeliverySimulator) progressPackage(ctx context.Context, pkg *entity.FulfillmentPackage) error {
	now := time.Now()

	switch pkg.Status {
	case entity.StatusPendingPickup:
		// Auto pickup sau khi scheduled time đã qua
		if pkg.PickupScheduledAt != nil && time.Now().After(*pkg.PickupScheduledAt) {
			pkg.Status = entity.StatusPickedUp
			pkg.PickupCompletedAt = &now
			pkg.CurrentHubLocation = strPtr("Đã lấy hàng từ người bán")
			pkg.LastScanAt = &now

			log.Printf("📦 Package %s: PENDING_PICKUP → PICKED_UP", pkg.PackageNumber)

			if err := s.repo.Update(ctx, pkg); err != nil {
				return err
			}
			return s.publisher.PublishPickedUp(ctx, pkg)
		}

	case entity.StatusPickedUp:
		// 1st location: Arrived at hub
		pkg.Status = entity.StatusAtHub
		pkg.CurrentHubLocation = strPtr("Kho trung chuyển - Đang phân loại")
		pkg.LastScanAt = &now

		log.Printf("📦 Package %s: PICKED_UP → AT_HUB", pkg.PackageNumber)

		if err := s.repo.Update(ctx, pkg); err != nil {
			return err
		}

	case entity.StatusAtHub:
		// 2nd location: In transit to destination hub
		pkg.Status = entity.StatusInTransit
		pkg.CurrentHubLocation = strPtr(s.getTransitLocation(pkg))
		pkg.LastScanAt = &now

		log.Printf("📦 Package %s: AT_HUB → IN_TRANSIT (%s)", pkg.PackageNumber, *pkg.CurrentHubLocation)

		if err := s.repo.Update(ctx, pkg); err != nil {
			return err
		}
		return s.publisher.PublishInTransit(ctx, pkg)

	case entity.StatusInTransit:
		// 3rd location: Out for delivery
		pkg.Status = entity.StatusOutForDelivery
		pkg.CurrentHubLocation = strPtr("Đang giao hàng đến bạn")
		pkg.LastScanAt = &now

		log.Printf("📦 Package %s: IN_TRANSIT → OUT_FOR_DELIVERY", pkg.PackageNumber)

		if err := s.repo.Update(ctx, pkg); err != nil {
			return err
		}
		return s.publisher.PublishOutForDelivery(ctx, pkg)

	case entity.StatusOutForDelivery:
		// 4th location: Delivered!
		pkg.Status = entity.StatusDelivered
		pkg.DeliveredAt = &now
		pkg.CurrentHubLocation = strPtr("Giao hàng thành công")
		pkg.LastScanAt = &now

		log.Printf("✅ Package %s: OUT_FOR_DELIVERY → DELIVERED", pkg.PackageNumber)

		if err := s.repo.Update(ctx, pkg); err != nil {
			return err
		}
		return s.publisher.PublishDelivered(ctx, pkg)
	}

	return nil
}

// getTransitLocation trả về location giả lập dựa trên delivery zone
func (s *AutoDeliverySimulator) getTransitLocation(pkg *entity.FulfillmentPackage) string {
	if pkg.DeliveryZone == nil {
		return "Đang vận chuyển"
	}

	switch *pkg.DeliveryZone {
	case "ZONE_NORTH":
		return "Kho Hà Nội - Đang vận chuyển"
	case "ZONE_SOUTH":
		return "Kho TP.HCM - Đang vận chuyển"
	case "ZONE_CENTRAL":
		return "Kho Đà Nẵng - Đang vận chuyển"
	default:
		return "Đang vận chuyển đến kho khu vực"
	}
}
