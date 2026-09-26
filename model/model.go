package model

type IModel[R any] interface {
	GetUniqueFields() string
	Fields() string	
	ToResponse() R
}

