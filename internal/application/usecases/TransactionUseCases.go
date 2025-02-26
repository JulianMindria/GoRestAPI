package usecases

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"
	"errors"
	"time"

	"github.com/google/uuid"
)

type TransactionUseCase struct {
	repo        repositories.TransactionRepository
	userRepo    repositories.UserRepository
	productRepo repositories.ProductRepository
	voucherRepo repositories.VoucherRepository
}

func NewTransactionUseCase(
	repo repositories.TransactionRepository,
	userRepo repositories.UserRepository,
	productRepo repositories.ProductRepository,
	voucherRepo repositories.VoucherRepository,
) *TransactionUseCase {
	return &TransactionUseCase{
		repo:        repo,
		userRepo:    userRepo,
		productRepo: productRepo,
		voucherRepo: voucherRepo,
	}
}

func (uc *TransactionUseCase) CreateTransaction(transaction *entities.Transaction) error {
	user, err := uc.userRepo.GetByID(transaction.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	product, err := uc.productRepo.GetByID(transaction.ProductID)
	if err != nil {
		return errors.New("product not found")
	}

	finalPrice := product.Price
	var cashback int

	if transaction.VoucherID != nil && *transaction.VoucherID != uuid.Nil {
		voucher, err := uc.voucherRepo.GetByID(*transaction.VoucherID)
		if err != nil {
			return errors.New("voucher not found")
		}

		if voucher.Expiration.Before(time.Now()) {
			return errors.New("voucher has expired")
		}

		if user.Points < voucher.CostInPoint {
			return errors.New("not enough points to use this voucher")
		}

		if voucher.Type == "discount" {
			finalPrice -= finalPrice * voucher.Value / 100
		} else if voucher.Type == "cashback" {
			cashback = finalPrice * voucher.Value / 100
		}

		user.Points -= voucher.CostInPoint
	}

	if user.Balance < finalPrice {
		return errors.New("insufficient balance")
	}

	user.Balance -= finalPrice
	user.Balance += cashback

	transaction.ID = uuid.New()
	transaction.CreatedAt = time.Now()
	transaction.Total = finalPrice

	err = uc.repo.Create(transaction)
	if err != nil {
		return errors.New("failed to create transaction")
	}

	err = uc.userRepo.Update(user)
	if err != nil {
		return errors.New("failed to update user balance and points")
	}

	return nil
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

func (uc *TransactionUseCase) GetTransactionDetailByUser(userID uuid.UUID) ([]entities.Transaction, error) {
	return uc.repo.GetTransactionDetailByUser(userID)
}
