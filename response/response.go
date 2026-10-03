package response

import "projeto-teste/model"

type IResponse[Model model.IModel[Model]] interface {
	FromModel(Model)
	GetName() string
}

type Pagination struct {
	TotalItems int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Data       interface{} `json:"data"`
}
