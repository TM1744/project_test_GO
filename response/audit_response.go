package response

import "projeto-teste/model"

type AuditResponse struct {
	ID     uint   `json:"id"`
	Date   string `json:"date"`
	Entity string `json:"entity"`
	Old    string `json:"old_record"`
	New    string `json:"new_record"`
	Action string `json:"action"`
}

func (a *AuditResponse) FromModel(m *model.Audit) {
	a.ID = m.ID
	a.Date = m.CreatedAt.String()
	a.Action = string(m.Action)
	a.Entity = string(m.Entity)
	a.New = m.New
	a.Old = m.Old
}

type AuditSimpleResponse struct {
	ID     uint   `json:"id"`
	Date   string `json:"date"`
	Entity string `json:"entity"`
	Action string `json:"action"`
}

func (a *AuditSimpleResponse) FromModel(m *model.Audit) {
	a.ID = m.ID
	a.Date = m.CreatedAt.String()
	a.Action = string(m.Action)
	a.Entity = string(m.Entity)
}
