package model

import (
	"time"

	"gorm.io/gorm"
)

type NaturalPerson struct {
	gorm.Model
	Name     *string `gorm:"type:varchar(150); not null"`
	CPF      *string `gorm:"type:varchar(11); not null; unique"`
	Birthday *time.Time
	PersonID *int64 `gorm:"not null; unique"`
}

func (a *NaturalPerson) GetUniqueFields() string {
	return "CPF and person ID"
}
