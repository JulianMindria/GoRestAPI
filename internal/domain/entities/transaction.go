package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UserID    uuid.UUID
	ProductID uuid.UUID
	VoucherID uuid.UUID
	Total     float64
}
