package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CustomBaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type User struct {
	CustomBaseModel
	Email    string
	Password string
}
