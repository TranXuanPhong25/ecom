package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Outbox struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	AggregateType string         `gorm:"column:aggregatetype;not null"`
	AggregateID   string         `gorm:"column:aggregateid;not null"`
	Type          string         `gorm:"column:type;not null"`
	Payload       datatypes.JSON `gorm:"column:payload;type:jsonb;not null"`
	Timestamp     time.Time      `gorm:"column:timestamp;not null"`
	TracingSpanID *string        `gorm:"column:tracingspanid"`
}
