package request

import "projeto-teste/model"

type IRequest[Model model.IModel] interface {
	Validate() error
	ToModel() Model
}
