package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Adresses *[]Adress `gorm:"foreignKey:PersonID"`
	Phone    *[]Phone  `gorm:"foreignKey:PersonID"`
	Email    *string   `gorm:"type:varchar(150);"`
}
