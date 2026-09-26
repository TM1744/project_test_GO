package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        *string `gorm:"varchar(70); not null; unique"`
	Description *string `gorm:"varchar(200); not null"`
}

func (a *Category) GetUniqueFields () string {
	return "name";
}
