package entities

import (
	"time"

	"github.com/google/uuid"
)

type Voucher struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	CostInPoint int       `json:"cost_in_point"`
	Expiration  time.Time `json:"expiration"`
	Type        string    `json:"type"`
	Value       float64   `json:"value"`
	BrandID     uuid.UUID `json:"brand_id"`
}
