package dtos

import "github.com/google/uuid"

type Shop struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Rating float64   `json:"rating"`
}

type GetShopsResponse struct {
	Shops       []Shop   `json:"shops"`
	NotFoundIDs []string `json:"notFoundIds,omitempty"`
}
