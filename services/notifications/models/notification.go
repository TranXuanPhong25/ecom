package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Message   string    `gorm:"type:text;not null"`
	Type      string    `gorm:"type:varchar(50);not null;default:'system'"` // order, promotion, system
	IsRead    bool      `gorm:"type:boolean;not null;default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
