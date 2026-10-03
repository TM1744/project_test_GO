package request

import "projeto-teste/model"

type IRequest[Model model.IModel[Model]] interface {
	Validate() error
	ToModel() Model
}