package models

import "gorm.io/gorm"

// Review represents a product review in the system
type Review struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	ProductID  uint   `json:"product_id" gorm:"not null"`
	UserID     uint   `json:"user_id" gorm:"not null"`
	Username   string `json:"username" gorm:"not null"`
	Rating     int    `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Title      string `json:"title" gorm:"not null;size:100"`
	Comment    string `json:"comment" gorm:"not null;type:text"`
}
