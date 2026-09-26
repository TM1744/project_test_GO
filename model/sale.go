package model

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	ClientID *uint       `gorm:"not null"`
	UserID   *uint       `gorm:"not null"`
	Value    *uint64     `gorm:"not null"`
	Items    *[]SaleItem `gorm:"foreignKey:SaleID"`
}

func (a *Sale) GetUniqueFields() string {
	return "none"
}
