package model

import "gorm.io/gorm"

type Subcategory struct {
	gorm.Model
	CategoryID  *int64  `gorm:"not null"`
	Name        *string `gorm:"varchar(70); not null; unique"`
	Description *string `gorm:"varchar(200); not null"`
}

func (a *Subcategory) GetUniqueFields () string {
	return "name";
}
