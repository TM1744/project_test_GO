package repository

import (
	"gorm.io/gorm"
	"projeto-teste/model"
)

type ProductRepository struct {
	*Repository[*model.Product]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Repository: InitRepository[*model.Product](
			db,
		),
	}
}
