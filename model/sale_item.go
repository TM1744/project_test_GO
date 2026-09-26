package model

import (
	"projeto-teste/enum"

	"gorm.io/gorm"
)

type SaleItem struct {
	gorm.Model
	ProductID       *int64                `gorm:"not null"`
	SaleID          *int64                `gorm:"not null"`
	Quantity        *int64                `gorm:"not null"`
	UnitPrice       *int64                `gorm:"not null"`
	MeasurementUnit *enum.MeasurementUnit `gorm:"type:varchar(12); not null"`
	ProductName     *string               `gorm:"not null"`
}

func (a *SaleItem) GetUniqueFields() string {
	return "none"
}
