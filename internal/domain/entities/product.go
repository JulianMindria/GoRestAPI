package entities

import "github.com/google/uuid"

type Product struct {
	ID      uuid.UUID
	Name    string
	Price   float64
	BrandID uuid.UUID
}
