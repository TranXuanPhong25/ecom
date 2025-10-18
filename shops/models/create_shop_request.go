package models

import "github.com/google/uuid"

type CreateShopRequest struct {
	Name         string    `json:"name" validate:"required,min=3,max=100"`
	OwnerID      uuid.UUID `json:"ownerId" validate:"required,uuid"`
	Location     string    `json:"location" validate:"required,min=3,max=255"`
	Logo         string    `json:"logo" validate:"omitempty,url"`
	Banner       string    `json:"banner" validate:"omitempty,url"`
	Email        string    `json:"email" validate:"required,email"`
	Phone        string    `json:"phone" validate:"omitempty,numeric"`
	BusinessType string    `json:"businessType" validate:"required,oneof=individual business"`
}
