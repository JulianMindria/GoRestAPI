package dto

import (
	"GoRestAPI/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Points  int       `json:"points"`
	Balance int       `json:"balance"`
}

type ProductDTO struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Price   int       `json:"price"`
	BrandID uuid.UUID `json:"brand_id"`
}

type VoucherDTO struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	CostInPoint int       `json:"cost_in_point"`
	Expiration  time.Time `json:"expiration"`
	Type        string    `json:"type"`
	Value       int       `json:"value"`
	BrandID     uuid.UUID `json:"brand_id"`
}

type TransactionDetailDTO struct {
	ID        uuid.UUID   `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UserID    uuid.UUID   `json:"user_id"`
	User      UserDTO     `json:"user"`
	ProductID uuid.UUID   `json:"product_id"`
	Product   ProductDTO  `json:"product"`
	VoucherID *uuid.UUID  `json:"voucher_id,omitempty"`
	Voucher   *VoucherDTO `json:"voucher,omitempty"`
	Total     int         `json:"total"`
}

func NewTransactionDetailDTO(transaction *entities.Transaction) TransactionDetailDTO {
	dto := TransactionDetailDTO{
		ID:        transaction.ID,
		CreatedAt: transaction.CreatedAt,
		UserID:    transaction.UserID,
		User: UserDTO{
			ID:      transaction.User.ID,
			Name:    transaction.User.Name,
			Points:  transaction.User.Points,
			Balance: transaction.User.Balance,
		},
		ProductID: transaction.ProductID,
		Product: ProductDTO{
			ID:      transaction.Product.ID,
			Name:    transaction.Product.Name,
			Price:   transaction.Product.Price,
			BrandID: transaction.Product.BrandID,
		},
		VoucherID: transaction.VoucherID,
		Total:     transaction.Total,
	}

	if transaction.Voucher != nil {
		dto.Voucher = &VoucherDTO{
			ID:          transaction.Voucher.ID,
			Code:        transaction.Voucher.Code,
			CostInPoint: transaction.Voucher.CostInPoint,
			Expiration:  transaction.Voucher.Expiration,
			Type:        transaction.Voucher.Type,
			Value:       transaction.Voucher.Value,
			BrandID:     transaction.Voucher.BrandID,
		}
	}

	return dto
}
