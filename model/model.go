package model

import (
	"projeto-teste/enum"
	"projeto-teste/fail"

	"gorm.io/gorm"
)

type IModel[Model any] interface {
	GetEntityType() enum.Entity
	GetUniqueFields() string
	ToString() string
	ValidateCreate() error
	OnCreate(*gorm.DB) error
	ValidateUpdate(Model) error
	MergeUpdate(Model)
	OnUpdate(*gorm.DB, Model) error
}

func auditCreation[Model IModel[Model]](db *gorm.DB, new Model) error {
	audit := &Audit{
		Action: enum.Create,
		Old:    new.ToString(),
		New:    new.ToString(),
		Entity: new.GetEntityType(),
	}

	err := db.Create(audit).Error

	if err != nil {
		return fail.StringToBadRequest("failed to generate audit record")
	}

	return nil
}

func auditUpdate[Model IModel[Model]](db *gorm.DB, old Model, new Model) error {
	audit := &Audit{
		Action: enum.Update,
		Old:    old.ToString(),
		New:    new.ToString(),
		Entity: new.GetEntityType(),
	}

	err := db.Create(audit).Error

	if err != nil {
		return fail.StringToBadRequest("failed to generate audit record")
	}

	return nil
}
