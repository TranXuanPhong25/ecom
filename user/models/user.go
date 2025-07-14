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
type User struct {
	CustomBaseModel
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}
