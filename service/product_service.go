package service

import (
	"projeto-teste/model"
	"projeto-teste/repository"
	"projeto-teste/request"
	"projeto-teste/response"
)

type ProductService struct {
	*Service[
		*model.Product,
		*request.ProductCreateRequest,
		*request.ProductUpdateRequest,
		*response.ProductResponse,
		*response.ProductSimpleResponse,
	]
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		Service: InitService[*model.Product, *request.ProductCreateRequest, *request.ProductUpdateRequest](
			repo.Repository,
			func() *model.Product { return &model.Product{} },
			func() []*model.Product { return make([]*model.Product, 0) },
			func() *response.ProductResponse { return &response.ProductResponse{} },
			func() *response.ProductSimpleResponse { return &response.ProductSimpleResponse{} },
			func() []*response.ProductSimpleResponse { return make([]*response.ProductSimpleResponse, 0) },
		),
	}
}
