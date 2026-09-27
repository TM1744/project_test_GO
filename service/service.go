package service

import (
	"projeto-teste/fail"
	"projeto-teste/model"
	"projeto-teste/repository"
	"projeto-teste/request"
	"projeto-teste/response"
	"projeto-teste/validator"
)

type Service[Model model.IModel, CRequest request.IRequest[Model], Response response.IResponse[Model]] struct {
	Repository  *repository.Repository[Model]
	newModel    func() Model
	newResponse func() Response
}

func InitService[
	Model model.IModel,
	CRequest request.IRequest[Model],
	Response response.IResponse[Model]](
	repository *repository.Repository[Model],
	newModel func() Model,
	newResponse func() Response) *Service[Model, CRequest, Response] {
	return &Service[Model, CRequest, Response]{
		Repository:  repository,
		newModel:    newModel,
		newResponse: newResponse,
	}
}

func (s *Service[Model, CRequest, Response]) Create(request CRequest) (Response, error) {
	var err error
	response := s.newResponse()
	entity := s.newModel()

	err = request.Validate()

	if err != nil {
		return response, fail.ErrorToBadRequest(err)
	}

	entity = request.ToModel()

	err = s.Repository.Create(entity)

	if err != nil {
		return response, err
	}

	response.FromModel(entity)

	return response, nil
}

func (s *Service[Model, CRequest, Response]) Read(id *int64) (Response, error) {
	var err error
	response := s.newResponse()

	err = validator.ValidIntNotLessThanZero(id, "ID")

	if err != nil {
		return response, fail.ErrorToBadRequest(err)
	}

	entity := s.newModel()

	err = s.Repository.Read(entity, id)

	if err != nil {
		return response, err
	}

	response.FromModel(entity)

	return response, nil
}

func (s *Service[Model, CRequest, Response]) Delete(id *int64) error {
	var err error
	model := s.newModel()

	err = validator.ValidIntNotLessThanZero(id, "ID")

	if err != nil {
		return err
	}

	err = s.Repository.Delete(model, id)

	if err != nil {
		return err
	}

	return nil
}
