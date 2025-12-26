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
type User struct {
	CustomBaseModel
	Email    string `gorm:"type:varchar(100);not null;uniqueIndex:idx_users_email"`
	Password string `gorm:"type:varchar(255);not null"`
}
