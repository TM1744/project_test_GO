package model

import (
	"projeto-teste/enum"

	"gorm.io/gorm"
)

type Audit struct {
	gorm.Model
	Entity enum.Entity
	Old    string
	New    string
	Action enum.Action
}

func (a *Audit) GetEntityType() enum.Entity      { return enum.Audit }
func (a *Audit) GetUniqueFields() string         { return "none" }
func (a *Audit) ToString() string                { return "" }
func (a *Audit) ValidateCreate() error           { return nil }
func (a *Audit) ValidateUpdate(*Audit) error     { return nil }
func (a *Audit) MergeUpdate(*Audit)              { return }
func (a *Audit) OnCreate(*gorm.DB) error         { return nil }
func (a *Audit) OnUpdate(*gorm.DB, *Audit) error { return nil }
