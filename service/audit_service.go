package service

import (
	"projeto-teste/model"
	"projeto-teste/repository"
	"projeto-teste/request"
	"projeto-teste/response"
)

type AuditService struct {
	*Service[
		*model.Audit,
		*request.AuditCreateRequest,
		*request.AuditUpdateRequest,
		*response.AuditResponse,
		*response.AuditSimpleResponse,
	]
}

func NewAuditService(repo *repository.AuditRepository) *AuditService {
	return &AuditService{
		Service: InitService[
			*model.Audit,
			*request.AuditCreateRequest,
			*request.AuditUpdateRequest,
		](
			repo.Repository,
			func() *model.Audit { return &model.Audit{} },
			func() []*model.Audit { return make([]*model.Audit, 0) },
			func() *response.AuditResponse { return &response.AuditResponse{} },
			func() *response.AuditSimpleResponse { return &response.AuditSimpleResponse{} },
			func() []*response.AuditSimpleResponse { return make([]*response.AuditSimpleResponse, 0) },
		),
	}
}
