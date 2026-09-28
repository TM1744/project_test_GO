package repository

import (
	"errors"
	"gorm.io/gorm"
	"projeto-teste/fail"
	"projeto-teste/model"
	"strings"
)

type Repository[Model model.IModel[Model]] struct {
	Database *gorm.DB
}

func InitRepository[Model model.IModel[Model]](database *gorm.DB) *Repository[Model] {
	return &Repository[Model]{Database: database}
}

func (r *Repository[Model]) Create(entity Model) error {
	err := r.Database.Create(entity).Error

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fail.StringToBadRequest(entity.GetUniqueFields() + " must be unique")
		} else {
			return fail.ErrorToInternalServerError(err)
		}
	}

	return nil
}

func (r *Repository[Model]) Read(entity Model, id *int64) error {
	err := r.Database.First(entity, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fail.StringToBadRequest("record not found. Check the ID")
		} else {
			return fail.ErrorToInternalServerError(err)
		}
	}

	return nil
}

func (r *Repository[Model]) Delete(entity Model, id *int64) error {

	err := r.Database.Delete(entity, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fail.StringToBadRequest("record not found. Check the ID")
		} else {
			return fail.ErrorToInternalServerError(err)
		}
	}

	return nil
}

func (r *Repository[Model]) Update(entity Model) error {
	err := r.Database.Save(entity).Error

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fail.StringToBadRequest(entity.GetUniqueFields() + " must be unique")
		} else {
			return fail.ErrorToInternalServerError(err)
		}
	}

	return nil
}
