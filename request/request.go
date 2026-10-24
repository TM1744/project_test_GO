package request

import "projeto-teste/model"

type IRequest[Model model.IModel[Model]] interface {
	Validate() error // Realiza a validação de campos independentes
	ToModel() Model // Transforma a requisição no seu respectiva entidade
}