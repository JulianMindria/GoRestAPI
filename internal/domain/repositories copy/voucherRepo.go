package repositories

import (
	"GoRestAPI/internal/domain/entities"  
	"github.com/google/uuid"
)

type VoucherRepository interface {
	Create(voucher *entities.Voucher) error
	GetByID(id uuid.UUID) (*entities.Voucher, error)
	GetAll() ([]entities.Voucher, error)
	Delete(id uuid.UUID) error
}
