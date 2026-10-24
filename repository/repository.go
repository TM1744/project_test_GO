package repository

import (
	"errors"
	"gorm.io/gorm"
	"projeto-teste/enum"
	"projeto-teste/fail"
	"projeto-teste/model"
	"strings"
)

type Repository[Model model.IModel[Model]] struct {
	Database *gorm.DB
}

func InitRepository[
	Model model.IModel[Model],
](
	database *gorm.DB,
) *Repository[Model] {
	return &Repository[Model]{
		Database: database,
	}
}

func (r *Repository[Model]) Create(entity Model) (err error) {
	if err = r.Database.Create(entity).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fail.StringToBadRequest(entity.GetUniqueFields() + " must be unique")
		} else {
			return fail.ErrorToInternalServerError(err)
		}
	}

	if err = entity.OnCreate(r.Database); err != nil {
		return err
	}

	if err = r.audit(enum.Create, entity, entity); err != nil {
		return err
	}

	return
}

func (r *Repository[Model]) ReadAll(totalItems *int64, limit *int, offset *int, entities *[]Model) (err error) {
	var entity Model

	if err = r.Database.Model(entity).Count(totalItems).Error; err != nil {
		return fail.StringToInternalServerError("failed to perform the total record count: " + err.Error())
	}

	if err = r.Database.Model(entity).Limit(*limit).Offset(*offset).Find(&entities).Error; err != nil {
		return fail.StringToInternalServerError("failed to retrieve all records")
	}

	return
}

func (r *Repository[Model]) Read(entity Model, id *int64) (err error) {
	if err = r.Database.First(entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fail.StringToBadRequest("record not found. Check the ID")
		} else {
			return fail.ErrorToInternalServerError(err)
		}
	}

	return
}

func (r *Repository[Model]) Delete(entity Model) (err error) {
	if err = r.Database.Delete(entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fail.StringToBadRequest("record not found. Check the ID")
		} else {
			return fail.ErrorToInternalServerError(err)
		}
	}

	if err = r.audit(enum.Delete, entity, entity); err != nil {
		return err
	}

	return
}

func (r *Repository[Model]) Update(entity Model, old Model) (err error) {
	if err = r.Database.Save(entity).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fail.StringToBadRequest(entity.GetUniqueFields() + " must be unique")
		} else {
			return fail.ErrorToInternalServerError(err)
		}
	}

	if err = entity.OnUpdate(r.Database, old); err != nil {
		return err
	}

	if err = r.audit(enum.Update, old, entity); err != nil {
		return err
	}

	return
}

func (r *Repository[Model]) audit(action enum.Action, oldEntity Model, newEntity Model) (err error) {
	audit := &model.Audit{
		Action: action,
		Old:    oldEntity.ToString(),
		New:    newEntity.ToString(),
		Entity: oldEntity.GetEntityType(),
	}

	if err = r.Database.Create(audit).Error; err != nil {
		return fail.StringToInternalServerError("failed to generate audit record")
	}

	return
}
