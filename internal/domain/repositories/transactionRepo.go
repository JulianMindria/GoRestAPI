package repositories

import (
	"GoRestAPI/internal/domain/entities"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
	GetByID(id uuid.UUID) (*entities.Transaction, error)
	GetAll() ([]entities.Transaction, error)
	Delete(id uuid.UUID) error
	GetTransactionDetailByUser(id uuid.UUID) ([]entities.Transaction, error)
}
