package model

import (
	"time"

	"gorm.io/gorm"
)

type LegalEntity struct {
	gorm.Model
	LegalName   *string `gorm:"type:varchar(150); not null"`
	TradeName   *string `gorm:"type:varchar(150); not null"`
	CNPJ        *string `gorm:"type:varchar(14); not null; unique"`
	OpeningDate *time.Time
	PersonID    *int64 `gorm:"not null; unique"`
}

func (a *LegalEntity) GetUniqueFields() string {
	return "CNPJ and person ID"
}
