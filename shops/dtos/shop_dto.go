package dtos

import (
	"github.com/google/uuid"
)

type ShopDTO struct {
	ID           uuid.UUID `json:"id"`
	OwnerID      uuid.UUID `json:"ownerId"`
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	Rating       float64   `json:"rating"`
	Logo         string    `json:"logo"`
	Banner       string    `json:"banner"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	BusinessType string    `json:"businessType"`
}

type GetShopsResponse struct {
	Shops       []ShopDTO `json:"shops"`
	NotFoundIDs []string  `json:"notFoundIds,omitempty"`
}
