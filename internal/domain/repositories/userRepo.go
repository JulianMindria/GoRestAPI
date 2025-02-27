package repositories

import (
	"GoRestAPI/internal/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(voucher *entities.User) error
	GetByID(id uuid.UUID) (*entities.User, error)
	GetAll() ([]entities.User, error)
	Delete(id uuid.UUID) error
	Update(user *entities.User) error
	UpdatePoints(id uuid.UUID, points int) error
	GetAvailableVouchersForUser(id uuid.UUID) ([]entities.Voucher, error)
}
