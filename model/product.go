package model

import (
	"projeto-teste/enum"
	"projeto-teste/fail"
	"projeto-teste/validator"

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

func (p *Product) Fields() string {
	return "ID: " + p.Name
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
	if p.ID != new.ID { return fail.StringToBadRequest("the parameter and request body IDs must be the same") }
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

// func (p *Product) ToString() (string, error) {
// 	id := converter.Int64ToString(int64(p.ID))
// 	category := converter.Int64ToString(*p.CategoryID)
// 	subcategory := converter.Int64ToString(*p.SubcategoryID)
// 	selling := converter.Int64ToString(p.SellingPrice)
// 	cost := converter.Int64ToString(p.CostPrice)
// 	stock := converter.Int64ToString(p.Stock)
// 	unit := p.Unit

// 	return ("ID: " + id +
// 		"\nCategory_ID: " + category +
// 		"\nSubcategory_ID: " + subcategory +
// 		"\nName: " + p.Name +
// 		"\nBrand: " + p.Brand +
// 		"\nDescription: " + p.Description +
// 		"\nSelling_Price: " + selling +
// 		"\nCost_Price: " + cost +
// 		"\nStock: " + stock +
// 		"\nUnit: " + p.Unit +
// 		"\nCreated_At: " + p.CreatedAt.String() +
// 		"\nUpdated_At: " + p.UpdatedAt.String() +
// 		"\nDeleted_At: " + p.DeletedAt.Time.String())

// }
