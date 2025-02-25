package usecases

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"
	"time"

	"github.com/google/uuid"
)

type VoucherUseCase struct {
	repo repositories.VoucherRepository
}

func NewVoucherUseCase(repo repositories.VoucherRepository) *VoucherUseCase {
	return &VoucherUseCase{repo: repo}
}

func (uc *VoucherUseCase) CreateVoucher(code string, costInPoint int, expiration time.Time, vType string, value float64, brandID uuid.UUID) (*entities.Voucher, error) {
	voucher := &entities.Voucher{
		ID:          uuid.New(),
		Code:        code,
		CostInPoint: costInPoint,
		Expiration:  expiration,
		Type:        vType,
		Value:       value,
		BrandID:     brandID,
	}
	err := uc.repo.Create(voucher)
	return voucher, err
}

func (uc *VoucherUseCase) GetVoucherByID(id uuid.UUID) (*entities.Voucher, error) {
	return uc.repo.GetByID(id)
}

func (uc *VoucherUseCase) GetAllVouchers() ([]entities.Voucher, error) {
	return uc.repo.GetAll()
}

func (uc *VoucherUseCase) DeleteVoucher(id uuid.UUID) error {
	return uc.repo.Delete(id)
}
