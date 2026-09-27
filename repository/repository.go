package repository

import (
	"errors"
	"projeto-teste/fail"
	"projeto-teste/model"
	"strings"
	"gorm.io/gorm"
)

type Repository[Model model.IModel] struct {
	Database *gorm.DB
}

func InitRepository[Model model.IModel](database *gorm.DB) Repository[Model] {
	return Repository[Model]{Database: database}
}

func (r *Repository[Model]) Create(entity Model) error {
	err := r.Database.Create(entity)

	if err != nil {
		if strings.Contains(err.Error.Error(), "UNIQUE constraint failed") {
			return fail.StringToBadRequest(entity.GetUniqueFields() + " must be unique")
		} else {
			return fail.ErrorToInternalServerError(err.Error)
		}
	}

	return nil
}

func (r *Repository[Model]) Read(entity Model, id *int64) error {
	err := r.Database.First(entity, id)

	if err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return fail.StringToBadRequest("record not found. Check the ID")
		} else {
			return fail.ErrorToInternalServerError(err.Error)
		}
	}

	return nil
}
