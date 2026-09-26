package repository

import (
	"errors"
	"projeto-teste/fail"
	"projeto-teste/model"
	"projeto-teste/response"

	"gorm.io/gorm"
)

type Repository[E model.IModel[R], R response.IResponse] struct {
	Database *gorm.DB
}

func InitRepository[E model.IModel[R], R response.IResponse](database *gorm.DB) Repository[E, R] {
	return Repository[E, R]{Database: database}
}

func (r *Repository[E, R]) Create(entity E) error {
	err := r.Database.Create(entity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fail.StringToBadRequest(entity.GetUniqueFields() + " must be unique")
		} else {
			return fail.ErrorToInternalServerError(err)
		}
	}

	return nil
}
