package repositories

import (
    "GoRestAPI/internal/domain/entities"
    "GoRestAPI/internal/domain/repositories"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type BrandRepository struct {
    db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) repositories.BrandRepository {
    return &BrandRepository{db: db}
}

func (r *BrandRepository) Create(brand *entities.Brand) error {
    brand.ID = uuid.New()
    return r.db.Create(brand).Error
}

func (r *BrandRepository) GetByID(id uuid.UUID) (*entities.Brand, error) {
    var brand entities.Brand
    err := r.db.First(&brand, "id = ?", id).Error
    return &brand, err
}

func (r *BrandRepository) GetAll() ([]entities.Brand, error) {
    var brands []entities.Brand
    err := r.db.Find(&brands).Error
    return brands, err
}

func (r *BrandRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&entities.Brand{}, "id = ?", id).Error
}
