package dtos

import "github.com/google/uuid"

type Shop struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Rating float64   `json:"rating"`
}
