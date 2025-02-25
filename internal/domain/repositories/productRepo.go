package repositories

import (
	"GoRestAPI/internal/domain/entities"

	"github.com/google/uuid"
)

type ProductRepository interface {
    Create(product *entities.Product) error
    GetByID(id uuid.UUID) (*entities.Product, error)
    GetAll() ([]entities.Product, error)
    Delete(id uuid.UUID) error
}
