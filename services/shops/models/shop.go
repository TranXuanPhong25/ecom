package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomBaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Shop struct {
	CustomBaseModel
	ProductsNum  int       `gorm:"type:int;default:0"`
	Follower     int       `gorm:"type:int;default:0"`
	Name         string    `gorm:"type:varchar(100);not null"`
	OwnerID      uuid.UUID `gorm:"type:uuid;unique;not null"`
	Location     string    `gorm:"type:varchar(255);not null"`
	Rating       float64   `gorm:"type:float;default:0"`
	Logo         string    `gorm:"type:varchar(255)"`
	Banner       string    `gorm:"type:varchar(255)"`
	Email        string    `gorm:"type:varchar(100);not null"`
	Phone        string    `gorm:"type:varchar(15)"`
	BusinessType string    `gorm:"type:varchar(20);default:'individual'"` // 'individual' or 'business'
}
