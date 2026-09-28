package service

import (
	"projeto-teste/fail"
	"projeto-teste/model"
	"projeto-teste/repository"
	"projeto-teste/request"
	"projeto-teste/response"
	"projeto-teste/validator"
)

type Service[Model model.IModel[Model], CRequest request.IRequest[Model], URequest request.IRequest[Model], Response response.IResponse[Model]] struct {
	Repository  *repository.Repository[Model]
	newModel    func() Model
	newResponse func() Response
}

func InitService[
	Model model.IModel[Model],
	CRequest request.IRequest[Model],
	URequest request.IRequest[Model],
	Response response.IResponse[Model]](
	repository *repository.Repository[Model],
	newModel func() Model,
	newResponse func() Response) *Service[Model, CRequest, URequest, Response] {
	return &Service[Model, CRequest, URequest, Response]{
		Repository:  repository,
		newModel:    newModel,
		newResponse: newResponse,
	}
}

func (s *Service[Model, CRequest, URequest, Response]) Create(request CRequest) (Response, error) {
	var err error
	response := s.newResponse()
	entity := s.newModel()

	err = request.Validate()

	if err != nil {
		return response, err
	}

	entity = request.ToModel()

	err = entity.ValidateCreate()

	if err != nil {
		return response, err
	}

	err = s.Repository.Create(entity)

	if err != nil {
		return response, err
	}

	response.FromModel(entity)

	return response, nil
}

func (s *Service[Model, CRequest, URequest, Response]) Read(id *int64) (Response, error) {
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

func (s *Service[Model, CRequest, URequest, Response]) Delete(id *int64) error {
	var err error
	entity := s.newModel()

	err = validator.ValidIntNotLessThanZero(id, "ID")

	if err != nil {
		return err
	}

	err = s.Repository.Delete(entity, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service[Model, CRequest, URequest, Response]) Update(request URequest, id *int64) (Response, error) {
	var err error
	response := s.newResponse()
	entity := s.newModel()
	oldRecord := s.newModel()

	err = validator.ValidIntNotLessThanZero(id, "ID")
	if err != nil {
		return response, err
	}

	err = request.Validate()
	if err != nil {
		return response, err
	}

	err = s.Repository.Read(oldRecord, id)
	if err != nil {
		return response, err
	}

	entity = request.ToModel()

	err = oldRecord.ValidateUpdate(entity)
	if err != nil {
		return response, err
	}

	entity.MergeUpdate(oldRecord)

	err = s.Repository.Update(entity)
	if err != nil {
		return response, err
	}

	response.FromModel(entity)

	return response, nil
}
