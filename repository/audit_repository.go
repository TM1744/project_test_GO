package repository

import (
	"projeto-teste/model"
	"gorm.io/gorm"
)

type AuditRepository struct {
	*Repository[*model.Audit]
}

func NewAuditRepository(db *gorm.DB) *AuditRepository {
	return &AuditRepository{
		Repository: InitRepository[*model.Audit](
			db,
		),
	}
}