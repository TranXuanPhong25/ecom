package models

import "github.com/google/uuid"

type CreateShopRequest struct {
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"required,min=10,max=500"`
	OwnerID     uuid.UUID `json:"ownerId" validate:"required,uuid"`
	Location    string    `json:"location" validate:"required,min=3,max=255"`
}
