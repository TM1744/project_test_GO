package request

import "projeto-teste/model"

type AuditCreateRequest struct {}

func (a *AuditCreateRequest) Validate() error { return nil }
func (a *AuditCreateRequest) ToModel() *model.Audit { return nil }

type AuditUpdateRequest struct {}

func (a *AuditUpdateRequest) Validate() error { return nil }
func (a *AuditUpdateRequest) ToModel() *model.Audit { return nil }


