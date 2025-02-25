package repositories

import (
	"GoRestAPI/internal/domain/entities"

	"github.com/google/uuid"
)

type BrandRepository interface {
    Create(brand *entities.Brand) error
    GetByID(id uuid.UUID) (*entities.Brand, error)
    GetAll() ([]entities.Brand, error)
    Delete(id uuid.UUID) error
}
