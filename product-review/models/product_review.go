package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID uint
	UserID    uint
	Rating    int
	Comment   string
}
