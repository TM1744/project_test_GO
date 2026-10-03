package model

type IModel [Model any] interface {
	GetUniqueFields() string
	GetName() string
	ValidateCreate() error
	ValidateUpdate(Model) error
	MergeUpdate(Model)
}
