package service

import (
	"projeto-teste/fail"
	"projeto-teste/model"
	"projeto-teste/repository"
	"projeto-teste/request"
	"projeto-teste/response"
)

type Service[RS response.IResponse, M model.IModel[RS], RQ request.IRequest[M]] struct {
	Repository *repository.Repository[M, RS]
}

func InitService[RS response.IResponse, M model.IModel[RS], RQ request.IRequest[M]](repository *repository.Repository[M, RS]) Service[RS, M, RQ] {
	return Service[RS, M, RQ]{Repository: repository}
}

func (s *Service[RS, M, RQ]) Create(request RQ) (RS, error) {
	var err error
	var response RS

	err = request.Validate()

	if err != nil {
		return response, fail.ErrorToBadRequest(err)
	}

	entity := request.ToModel()

	err = s.Repository.Create(entity)

	if err != nil {
		return response, err
	}

	response = entity.ToResponse()

	return response, nil
}
