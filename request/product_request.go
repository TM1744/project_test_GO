package request

import (
	"errors"
	"projeto-teste/enum"
	"projeto-teste/model"
	"projeto-teste/validator"
)

type ProductRequest struct {
	CategoryID    *int64 `json:"category_id"`
	SubcategoryID *int64 `json:"subcategory_id"`
	Description   string `json:"description"`
	SellingPrice  int64  `json:"selling_price"`
	CostPrice     int64  `json:"cost_price"`
	Stock         int64  `json:"stock"`
}

type ProductCreateRequest struct {
	ProductRequest
	Name  string               `json:"name"`
	Brand string               `json:"brand"`
	Unit  enum.MeasurementUnit `json:"unit"`
}

func (p *ProductCreateRequest) Validate() error {

	name := "product name"
	brand := "product brand"

	if err := validator.ValidStrintIsNotBlank(&p.Name, name); err != nil {
		return err
	}
	if err := validator.ValidStringMaxLength(&p.Name, 150, name); err != nil {
		return err
	}
	if err := validator.ValidStringMaxLength(&p.Brand, 70, brand); err != nil {
		return err
	}

	if !p.Unit.IsValid() {
		return errors.New("measurement unit type is invalid")
	}

	validateProductCommonFields(&p.ProductRequest)

	return nil
}

func (p *ProductCreateRequest) ToModel() *model.Product {
	return &model.Product{
		CategoryID:    p.CategoryID,
		SubcategoryID: p.SubcategoryID,
		Name:          p.Name,
		Brand:         p.Brand,
		Description:   p.Description,
		SellingPrice:  p.SellingPrice,
		CostPrice:     p.CostPrice,
		Stock:         p.Stock,
		Unit:          p.Unit,
	}
}

type ProductUpdateRequest struct {
	ProductRequest
	ID uint `json:"id"`
}

func (p *ProductUpdateRequest) Validate() error {

	id := "product ID"
	n := int64(p.ID)

	if err := validator.ValidIntHighThanZero(&n, id); err != nil {
		return err
	}

	validateProductCommonFields(&p.ProductRequest)

	return nil
}

func (p *ProductUpdateRequest) ToModel() *model.Product {
	return &model.Product{
		ID:            p.ID,
		CategoryID:    p.CategoryID,
		SubcategoryID: p.SubcategoryID,
		Description:   p.Description,
		SellingPrice:  p.SellingPrice,
		CostPrice:     p.CostPrice,
		Stock:         p.Stock,
	}
}

func validateProductCommonFields(p *ProductRequest) error {
	category := "category ID"
	subcategory := "subcategory ID"
	description := "product description"
	price := "product selling price"
	cost := "product cost price"
	stock := "product stock"

	if err := validator.ValidIntHighThanZeroButCanNil(p.CategoryID, category); err != nil {
		return err
	}
	if err := validator.ValidIntHighThanZeroButCanNil(p.SubcategoryID, subcategory); err != nil {
		return err
	}
	if err := validator.ValidStringMaxLength(&p.Description, 200, description); err != nil {
		return err
	}
	if err := validator.ValidIntHighThanZero(&p.SellingPrice, price); err != nil {
		return err
	}
	if err := validator.ValidIntNotLessThanZero(&p.CostPrice, cost); err != nil {
		return err
	}
	if err := validator.ValidIntNotLessThanZero(&p.Stock, stock); err != nil {
		return err
	}

	return nil
}
