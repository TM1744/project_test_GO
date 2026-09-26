package model

import (
	"projeto-teste/enum"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	PersonType enum.PersonType `gorm:"type:varchar(20); not null"`
	PersonId   uint            `gorm:"not null; unique"`
}

func (a *Client) GetUniqueFields () string {
	return "person ID";
}