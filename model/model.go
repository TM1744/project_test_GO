package model

type IModel interface {
	GetUniqueFields() string
	Fields() string	
}

