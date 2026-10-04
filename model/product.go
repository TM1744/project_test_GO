package model

import (
	"projeto-teste/converter"
	"projeto-teste/enum"
	"projeto-teste/fail"
	"projeto-teste/validator"
	"strconv"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryID    *int64
	SubcategoryID *int64
	Name          string               `gorm:"type:varchar(150); not null; unique"`
	Brand         string               `gorm:"type:varchar(70);"`
	Description   string               `gorm:"type:varchar(200);"`
	SellingPrice  int64                `gorm:"not null"`
	CostPrice     int64                `gorm:"not null"`
	Stock         int64                `gorm:"not null"`
	Unit          enum.MeasurementUnit `gorm:"type:varchar(12); not null"`
}

func (p *Product) GetUniqueFields() string {
	return "name"
}

func (p *Product) GetEntityType() enum.Entity {
	return enum.Product
}

func (p *Product) ToString() string {
	return ("ID: " + strconv.Itoa(int(p.ID)) +
		"\nCreated_At: " + p.CreatedAt.String() +
		"\nUpdated_At: " + p.UpdatedAt.String() +
		"\nDeleted_At: " + p.DeletedAt.Time.String() +
		"\nCategory_ID: " + converter.PInt64ToString(p.CategoryID) +
		"\nSubcategory_ID: " + converter.PInt64ToString(p.SubcategoryID) +
		"\nName: " + p.Name +
		"\nBrand: " + p.Brand +
		"\nDescription: " + p.Description +
		"\nSelling_Price: " + converter.Int64ToString(p.SellingPrice) +
		"\nCost_Price: " + converter.Int64ToString(p.CostPrice) +
		"\nStock: " + converter.Int64ToString(p.Stock) +
		"\nUnit: " + string(p.Unit))
}

func (p *Product) ValidateCreate() error {
	price := "product selling price"
	cost := "product cost price"

	if err := validator.ValidIntIsHighestThanValue(&p.SellingPrice, &p.CostPrice, price, cost); err != nil {
		return err
	}

	return nil
}

func (p *Product) ValidateUpdate(new *Product) error {
	if p.ID != new.ID {
		return fail.StringToBadRequest("the parameter and request body IDs must be the same")
	}
	return nil
}

func (p *Product) MergeUpdate(old *Product) {
	p.Name = old.Name
	p.Brand = old.Brand
	p.CreatedAt = old.CreatedAt
	p.DeletedAt = old.DeletedAt
	p.ID = old.ID
	p.Unit = old.Unit
}

func (p *Product) OnCreate(db *gorm.DB) error {
	if err := auditCreation(db, p); err != nil {
		return err
	}

	return nil
}

func (p *Product) OnUpdate(db *gorm.DB, old *Product) error {
	if err := auditUpdate(db, old, p); err != nil {
		return err
	}

	return nil
}
