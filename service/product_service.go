package service

import (
	"projeto-teste/model"
	"projeto-teste/repository"
	"projeto-teste/request"
	"projeto-teste/response"
)

type ProductService struct {
	*Service[*model.Product, *request.ProductCreateRequest, *request.ProductUpdateRequest, *response.ProductResponse]
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		Service: InitService[*model.Product, *request.ProductCreateRequest, *request.ProductUpdateRequest](
			repo.Repository,
			func() *model.Product { return &model.Product{} },
			func() *response.ProductResponse { return &response.ProductResponse{} },
		),
	}
}
