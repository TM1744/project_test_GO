package service

import (
	"projeto-teste/model"
	"projeto-teste/repository"
	"projeto-teste/request"
	"projeto-teste/response"
)

type ProductService struct {
	Service[response.ProductResponse, *model.Product, *request.ProductRequest]
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{
        Service: InitService[response.ProductResponse, *model.Product, *request.ProductRequest](&repo.Repository),
    }
}
