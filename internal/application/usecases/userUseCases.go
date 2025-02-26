package usecases

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserUseCase struct {
	repo        repositories.UserRepository
	voucherRepo repositories.VoucherRepository
}

func NewUserUseCase(repo repositories.UserRepository, voucherRepo repositories.VoucherRepository) *UserUseCase {
	return &UserUseCase{
		repo:        repo,
		voucherRepo: voucherRepo,
	}
}

func (uc *UserUseCase) CreateUser(name string, Point int, Money int) (*entities.User, error) {
	user := &entities.User{
		ID:      uuid.New(),
		Name:    name,
		Points:  Point,
		Balance: Money,
	}
	err := uc.repo.Create(user)
	return user, err
}

func (uc *UserUseCase) GetUserByID(id uuid.UUID) (*entities.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUseCase) GetAllUsers() ([]entities.User, error) {
	return uc.repo.GetAll()
}

func (uc *UserUseCase) DeleteVoucher(id uuid.UUID) error {
	return uc.repo.Delete(id)
}

func (uc *UserUseCase) UpdatePoints(id uuid.UUID, points int) error {
	return uc.repo.UpdatePoints(id, points)
}

func (uc *UserUseCase) UpdateUser(user *entities.User) error {
	existingUser, err := uc.repo.GetByID(user.ID)
	if err != nil {
		return err
	}

	existingUser.Name = user.Name
	existingUser.Points = user.Points
	existingUser.Balance = user.Balance

	return uc.repo.Update(existingUser)
}

func (uc *UserUseCase) GetAvailableVouchersForUser(userID uuid.UUID) ([]entities.Voucher, error) {
	user, err := uc.repo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	vouchers, err := uc.voucherRepo.GetAll()
	if err != nil {
		return nil, errors.New("failed to retrieve vouchers")
	}

	availableVouchers := []entities.Voucher{}
	for _, voucher := range vouchers {
		if voucher.Expiration.After(time.Now()) && user.Points >= voucher.CostInPoint {
			availableVouchers = append(availableVouchers, voucher)
		}
	}

	return availableVouchers, nil
}
