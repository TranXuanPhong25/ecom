package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CustomBaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Shop struct {
	CustomBaseModel
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text;not null"`
	OwnerID     uuid.UUID `gorm:"type:uuid;unique;not null"`
	Location    string    `gorm:"type:varchar(255);not null"`
	Rating      float64   `gorm:"type:float;default:0"`
}
