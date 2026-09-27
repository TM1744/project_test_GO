package response

import "projeto-teste/model"

type IResponse[Model model.IModel] interface {
	FromModel(Model)
}