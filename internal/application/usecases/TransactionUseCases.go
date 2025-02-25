package usecases

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"

	"github.com/google/uuid"
)

type TransactionUseCase struct {
	repo repositories.TransactionRepository
}

func NewTransactionUseCase(repo repositories.TransactionRepository) *TransactionUseCase {
	return &TransactionUseCase{repo: repo}
}

func (uc *TransactionUseCase) CreateTransaction(transaction *entities.Transaction) error {
	return uc.repo.Create(transaction)
}

func (uc *TransactionUseCase) GetTransactionByID(id uuid.UUID) (*entities.Transaction, error) {
	return uc.repo.GetByID(id)
}

func (uc *TransactionUseCase) GetAllTransactions() ([]entities.Transaction, error) {
	return uc.repo.GetAll()
}

func (uc *TransactionUseCase) DeleteTransaction(id uuid.UUID) error {
	return uc.repo.Delete(id)
}
