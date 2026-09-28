package model

type IModel [Model any] interface {
	GetUniqueFields() string
	ValidateCreate() error
	ValidateUpdate(Model) error
	MergeUpdate(Model)
}
