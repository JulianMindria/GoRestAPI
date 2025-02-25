package usecases

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"

	"github.com/google/uuid"
)

type BrandUseCase struct {
	repo repositories.BrandRepository
}

func NewBrandUseCase(repo repositories.BrandRepository) *BrandUseCase {
	return &BrandUseCase{repo: repo}
}

func (uc *BrandUseCase) CreateBrand(brand *entities.Brand) error {
	return uc.repo.Create(brand)
}

func (uc *BrandUseCase) GetBrandByID(id uuid.UUID) (*entities.Brand, error) {
	return uc.repo.GetByID(id)
}

func (uc *BrandUseCase) GetAllBrands() ([]entities.Brand, error) {
	return uc.repo.GetAll()
}

func (uc *BrandUseCase) DeleteBrand(id uuid.UUID) error {
	return uc.repo.Delete(id)
}
