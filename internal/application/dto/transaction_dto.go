package dto

import (
	"GoRestAPI/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type TransactionDTO struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UserID    uuid.UUID  `json:"user_id"`
	ProductID uuid.UUID  `json:"product_id"`
	VoucherID *uuid.UUID `json:"voucher_id,omitempty"`
	Total     int        `json:"total"`
}

func NewTransactionDTO(transaction *entities.Transaction) TransactionDTO {
	return TransactionDTO{
		ID:        transaction.ID,
		CreatedAt: transaction.CreatedAt,
		UserID:    transaction.UserID,
		ProductID: transaction.ProductID,
		VoucherID: transaction.VoucherID,
		Total:     transaction.Total,
	}
}
