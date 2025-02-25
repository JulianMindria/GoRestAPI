package repositories

import (
    "GoRestAPI/internal/domain/entities"
    "GoRestAPI/internal/domain/repositories"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repositories.ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *entities.Product) error {
    product.ID = uuid.New()
    return r.db.Create(product).Error
}

func (r *ProductRepository) GetByID(id uuid.UUID) (*entities.Product, error) {
    var product entities.Product
    err := r.db.First(&product, "id = ?", id).Error
    return &product, err
}

func (r *ProductRepository) GetAll() ([]entities.Product, error) {
    var products []entities.Product
    err := r.db.Find(&products).Error
    return products, err
}

func (r *ProductRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&entities.Product{}, "id = ?", id).Error
}
