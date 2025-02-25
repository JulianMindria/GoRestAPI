package repositories

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(voucher *entities.User) error {
	voucher.ID = uuid.New()
	return r.db.Create(voucher).Error
}

func (r *UserRepository) GetByID(id uuid.UUID) (*entities.User, error) {
	var voucher entities.User
	err := r.db.First(&voucher, "id = ?", id).Error
	return &voucher, err
}

func (r *UserRepository) GetAll() ([]entities.User, error) {
	var vouchers []entities.User
	err := r.db.Find(&vouchers).Error
	return vouchers, err
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.User{}, "id = ?", id).Error
}

func (r *UserRepository) UpdatePoints(id uuid.UUID, points int) error {
	// Use GORM's Model and Updates to update the points field
	return r.db.Model(&entities.User{}).Where("id = ?", id).Update("points", points).Error
}
