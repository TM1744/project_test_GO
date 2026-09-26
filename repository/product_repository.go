package repository

import (
	"projeto-teste/model"
	"projeto-teste/response"

	"gorm.io/gorm"
)


type ProductRepository struct {
	Repository[*model.Product, response.ProductResponse]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{
        Repository: InitRepository[*model.Product](db),
    }
}