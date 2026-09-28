package handler

import (
	"projeto-teste/model"
	"projeto-teste/request"
	"projeto-teste/response"
	"projeto-teste/service"
)

type ProductHandler struct {
	*Handler[*model.Product, *request.ProductCreateRequest, *request.ProductUpdateRequest, *response.ProductResponse]
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		Handler: InitHandler(
			service.Service,
			func() *request.ProductCreateRequest { return &request.ProductCreateRequest{} },
			func() *request.ProductUpdateRequest { return &request.ProductUpdateRequest{} },
		),
	}
}
