package model

import (
	"projeto-teste/enum"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PersonType enum.PersonType `gorm:"type:varchar(20); not null"`
	PersonId   uint            `gorm:"not null; unique"`
	Role       enum.Role       `gorm:"type:varchar(20); not null"`
	Salary     uint64          `gorm:"not null"`
}

func (a *User) GetUniqueFields () string {
	return "person ID";
}
