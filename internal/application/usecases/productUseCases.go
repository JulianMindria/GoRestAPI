package usecases

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"

	"github.com/google/uuid"
)

type ProductUseCase struct {
	repo repositories.ProductRepository
}

func NewProductUseCase(repo repositories.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) CreateProduct(product *entities.Product) error {
	return uc.repo.Create(product)
}

func (uc *ProductUseCase) GetProductByID(id uuid.UUID) (*entities.Product, error) {
	return uc.repo.GetByID(id)
}

func (uc *ProductUseCase) GetAllProducts() ([]entities.Product, error) {
	return uc.repo.GetAll()
}

func (uc *ProductUseCase) DeleteProduct(id uuid.UUID) error {
	return uc.repo.Delete(id)
}
