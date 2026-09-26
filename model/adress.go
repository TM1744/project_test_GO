package model

import (
	"gorm.io/gorm"
)

type Adress struct {
	gorm.Model
	AdressName   *string          `gorm:"type:varchar(100)"`
	AdressNumber *string          `gorm:"type:varchar(20)"`
	CEP          *string          `gorm:"type:varchar(8)"`
	Complement   *string          `gorm:"type:varchar(20)"`
	Neighborhood *string          `gorm:"type:varchar(50)"`
	City         *string          `gorm:"type:varchar(50)"`
	State        *string          `gorm:"type:varchar(2)"`
	PersonID     *uint            `gorm:"not null"`
}

func (a *Adress) GetUniqueFields() string {
	return "none"
}
