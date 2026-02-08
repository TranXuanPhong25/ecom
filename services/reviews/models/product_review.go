package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Review represents a product review in the system
type Review struct {
	gorm.Model           // ID, CreatedAt, UpdatedAt, DeletedAt
	ProductID  uint      `json:"productId" gorm:"not null"`
	UserID     uuid.UUID `json:"userId" gorm:"type:uuid;not null"`
	Username   string    `json:"username" gorm:"not null"`
	Rating     int       `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment    string    `json:"comment" gorm:"not null;type:text"`
}
