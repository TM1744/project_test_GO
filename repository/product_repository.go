package repository

import (
	"projeto-teste/model"
	"gorm.io/gorm"
)


type ProductRepository struct {
	Repository[*model.Product]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{
        Repository: InitRepository[*model.Product](db),
    }
}