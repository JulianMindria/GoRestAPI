package repositories

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repositories.TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction *entities.Transaction) error {
	transaction.ID = uuid.New()
	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) GetByID(id uuid.UUID) (*entities.Transaction, error) {
	var transaction entities.Transaction
	err := r.db.First(&transaction, "id = ?", id).Error
	return &transaction, err
}

func (r *TransactionRepository) GetAll() ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Transaction{}, "id = ?", id).Error
}

func (r *TransactionRepository) GetTransactionDetailByUser(userID uuid.UUID) ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	err := r.db.Debug().
		Preload("User").
		Preload("Product").
		Preload("Voucher").
		Where("user_id = ?", userID).
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
