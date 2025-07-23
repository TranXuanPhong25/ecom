package dtos

import (
	"github.com/google/uuid"
)

type ShopDTO struct {
	ID          uuid.UUID `json:"id"`
	OwnerID     uuid.UUID `json:"ownerId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Rating      float64   `json:"rating"`
}
