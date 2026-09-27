package handler

import (
	"projeto-teste/model"
	"projeto-teste/request"
	"projeto-teste/response"
	"projeto-teste/service"
)

type ProductHandler struct {
	Handler[*model.Product, *request.ProductRequest, *response.ProductResponse]
}

func NewProductHandler (service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		Handler: InitHandler(
			&service.Service,
			func () *request.ProductRequest { return &request.ProductRequest{} },
		),
	}
}