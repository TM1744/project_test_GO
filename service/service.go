package service

import (
	"math"
	"projeto-teste/fail"
	"projeto-teste/model"
	"projeto-teste/repository"
	"projeto-teste/request"
	"projeto-teste/response"
	"projeto-teste/validator"
)

type Service[
	Model model.IModel[Model],
	CRequest request.IRequest[Model],
	URequest request.IRequest[Model],
	Response response.IResponse[Model],
	SResponse response.IResponse[Model],
] struct {
	Repository       *repository.Repository[Model]
	newModel         func() Model
	newModelList     func() []Model
	newResponse      func() Response
	newSResponse     func() SResponse
	newSResponseList func() []SResponse
}

func InitService[
	Model model.IModel[Model],
	CRequest request.IRequest[Model],
	URequest request.IRequest[Model],
	Response response.IResponse[Model],
	SResponse response.IResponse[Model],
](
	repository *repository.Repository[Model],
	newModel func() Model,
	newModelList func() []Model,
	newResponse func() Response,
	newSResponse func() SResponse,
	newSResponseList func() []SResponse,
) *Service[
	Model,
	CRequest,
	URequest,
	Response,
	SResponse,
] {
	return &Service[
		Model,
		CRequest,
		URequest,
		Response,
		SResponse,
	]{
		Repository:       repository,
		newModel:         newModel,
		newModelList:     newModelList,
		newResponse:      newResponse,
		newSResponse:     newSResponse,
		newSResponseList: newSResponseList,
	}
}

func (s *Service[Model, CRequest, URequest, Response, SResponse]) Create(request CRequest) (Response, error) {
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

func (s *Service[Model, CRequest, URequest, Response, SResponse]) ReadAll(page *int, limit *int) (*response.Pagination, error) {
	pagination := &response.Pagination{}

	if *page < 1 {
		*page = 1
	}

	if *limit < 1 || *limit > 100 {
		*limit = 10
	}

	offset := (*page - 1) * *limit

	var totalItems int64
	entities := s.newModelList()

	err := s.Repository.ReadAll(&totalItems, limit, &offset, &entities)

	if err != nil {
		return pagination, err
	}

	responses := s.newSResponseList()

	for _, entity := range entities {
		response := s.newSResponse()
		response.FromModel(entity)
		responses = append(responses, response)
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(*limit)))

	pagination.TotalItems = totalItems
	pagination.TotalPages = totalPages
	pagination.Page = *page
	pagination.Limit = *limit
	pagination.Data = &responses

	return pagination, nil
}

func (s *Service[Model, CRequest, URequest, Response, SResponse]) Read(id *int64) (Response, error) {
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

func (s *Service[Model, CRequest, URequest, Response, SResponse]) Delete(id *int64) error {
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

func (s *Service[Model, CRequest, URequest, Response, SResponse]) Update(request URequest, id *int64) (Response, error) {
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

	err = s.Repository.Update(entity, oldRecord)
	if err != nil {
		return response, err
	}

	response.FromModel(entity)

	return response, nil
}
