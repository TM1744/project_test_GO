package handler

import (
	"projeto-teste/model"
	"projeto-teste/request"
	"projeto-teste/response"
	"projeto-teste/service"
)

type AuditHandler struct {
	*Handler[
		*model.Audit,
		*request.AuditCreateRequest,
		*request.AuditUpdateRequest,
		*response.AuditResponse,
		*response.AuditSimpleResponse,
	]
}

func NewAuditHandler(service *service.AuditService) *AuditHandler {
	return &AuditHandler{
		Handler: InitHandler(
			service.Service,
			func() *request.AuditCreateRequest { return &request.AuditCreateRequest{} },
			func() *request.AuditUpdateRequest { return &request.AuditUpdateRequest{} },
		),
	}
}