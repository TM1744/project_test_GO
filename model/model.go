package model

type IModel interface {
	GetUniqueFields() string
	Validate() error
}

