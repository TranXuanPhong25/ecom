package entity

import "time"

type PackageStatus string

const (
	StatusPendingPickup    PackageStatus = "PENDING_PICKUP"
	StatusPickedUp         PackageStatus = "PICKED_UP"
	StatusAtHub            PackageStatus = "AT_HUB"
	StatusInTransit        PackageStatus = "IN_TRANSIT"
	StatusOutForDelivery   PackageStatus = "OUT_FOR_DELIVERY"
	StatusDelivered        PackageStatus = "DELIVERED"
	StatusDeliveryFailed   PackageStatus = "DELIVERY_FAILED"
	StatusReturnedToSeller PackageStatus = "RETURNED_TO_SELLER"
)

type FulfillmentPackage struct {
	ID            int64  `gorm:"primaryKey;column:id"`
	PackageNumber string `gorm:"uniqueIndex;column:package_number;type:varchar(50);not null"`
	OrderID       int64  `gorm:"column:order_id;not null;index"`

	// Seller/pickup info
	ShopID             string     `gorm:"column:shop_id;type:varchar(50);not null;index"`
	PickupAddress      string     `gorm:"column:pickup_address;type:text;not null"`
	PickupContactName  *string    `gorm:"column:pickup_contact_name;type:varchar(100)"`
	PickupContactPhone *string    `gorm:"column:pickup_contact_phone;type:varchar(20)"`
	PickupScheduledAt  *time.Time `gorm:"column:pickup_scheduled_at"`
	PickupCompletedAt  *time.Time `gorm:"column:pickup_completed_at"`

	// Package status
	Status PackageStatus `gorm:"column:status;type:varchar(50);not null;default:'PENDING_PICKUP';index"`

	// Transit tracking
	CurrentHubLocation *string    `gorm:"column:current_hub_location;type:varchar(100)"`
	LastScanAt         *time.Time `gorm:"column:last_scan_at"`
	EstimatedDelivery  *time.Time `gorm:"column:estimated_delivery"`

	// Delivery info
	DeliveryAddress      string  `gorm:"column:delivery_address;type:text;not null"`
	DeliveryContactName  string  `gorm:"column:delivery_contact_name;type:varchar(100);not null"`
	DeliveryContactPhone string  `gorm:"column:delivery_contact_phone;type:varchar(20);not null"`
	DeliveryZone         *string `gorm:"column:delivery_zone;type:varchar(50);index"`
	DeliveryPartner      *string `gorm:"column:delivery_partner;type:varchar(100)"`
	TrackingNumber       *string `gorm:"column:tracking_number;type:varchar(100)"`

	// Delivery attempts
	DeliveryAttempts      int        `gorm:"column:delivery_attempts;default:0"`
	LastDeliveryAttemptAt *time.Time `gorm:"column:last_delivery_attempt_at"`
	DeliveryFailureReason *string    `gorm:"column:delivery_failure_reason;type:text"`
	DeliveredAt           *time.Time `gorm:"column:delivered_at"`
	DeliverySignatureURL  *string    `gorm:"column:delivery_signature_url;type:text"`

	// Package details
	WeightGrams         *int            `gorm:"column:weight_grams"`
	Dimensions          *map[string]any `gorm:"column:dimensions;type:jsonb;serializer:json"`
	SpecialInstructions *string         `gorm:"column:special_instructions;type:text"`

	// Metadata
	Metadata *map[string]any `gorm:"column:metadata;type:jsonb;serializer:json"`

	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:now()"`
}

func (FulfillmentPackage) TableName() string {
	return "fulfillment_packages"
}
