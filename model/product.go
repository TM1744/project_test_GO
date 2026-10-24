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
		" | Created_At: " + p.CreatedAt.String() +
		" | Updated_At: " + p.UpdatedAt.String() +
		" | Deleted_At: " + p.DeletedAt.Time.String() +
		" | Category_ID: " + converter.PInt64ToString(p.CategoryID) +
		" | Subcategory_ID: " + converter.PInt64ToString(p.SubcategoryID) +
		" | Name: " + p.Name +
		" | Brand: " + p.Brand +
		" | Description: " + p.Description +
		" | Selling_Price: " + converter.Int64ToString(p.SellingPrice) +
		" | Cost_Price: " + converter.Int64ToString(p.CostPrice) +
		" | Stock: " + converter.Int64ToString(p.Stock) +
		" | Unit: " + string(p.Unit))
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
	return nil
}
 
func (p *Product) OnUpdate(db *gorm.DB, old *Product) error {
	return nil
}
