package handler

import (
	"projeto-teste/model"
	"projeto-teste/request"
	"projeto-teste/response"
	"projeto-teste/service"
)

type ProductHandler struct {
	Handler[response.ProductResponse, *model.Product, *request.ProductRequest]
}

func NewProductHandler (service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		Handler: InitHandler(&service.Service),
	}
}