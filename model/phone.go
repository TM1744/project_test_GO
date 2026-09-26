package model

import (
	"gorm.io/gorm"
)

type Phone struct {
	gorm.Model
	DDD        *string          `gorm:"type:varchar(2); not null"`
	Number     *string          `gorm:"type:varchar(9); not null"`
	PersonID   *uint            `gorm:"not null"`
}

func (a *Phone) GetUniqueFields() string {
	return "none"
}
