package models

import (
	"time"

	"github.com/google/uuid"
)

type ConversationType string
type ConversationStatus string

const (
	ConversationTypeCustomerShop   ConversationType = "customer_shop"
	ConversationTypeCustomerBot    ConversationType = "customer_bot"
	ConversationTypeCustomerSystem ConversationType = "customer_system"

	ConversationStatusOpen     ConversationStatus = "open"
	ConversationStatusResolved ConversationStatus = "resolved"
	ConversationStatusPending  ConversationStatus = "pending"
)

type Conversation struct {
	ID        uuid.UUID          `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Type      ConversationType   `gorm:"type:varchar(20);not null" json:"type"`
	Status    ConversationStatus `gorm:"type:varchar(20);not null;default:'open'" json:"status"`
	CreatedAt time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
}
