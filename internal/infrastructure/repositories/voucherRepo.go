package repositories

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) repositories.VoucherRepository {
	return &VoucherRepository{db: db}
}

func (r *VoucherRepository) Create(voucher *entities.Voucher) error {
	voucher.ID = uuid.New()
	return r.db.Create(voucher).Error
}

func (r *VoucherRepository) GetByID(id uuid.UUID) (*entities.Voucher, error) {
	var voucher entities.Voucher
	err := r.db.First(&voucher, "id = ?", id).Error
	return &voucher, err
}

func (r *VoucherRepository) GetAll() ([]entities.Voucher, error) {
	var vouchers []entities.Voucher
	err := r.db.Find(&vouchers).Error
	return vouchers, err
}

func (r *VoucherRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Voucher{}, "id = ?", id).Error
}
