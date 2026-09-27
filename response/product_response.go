package response

import (
	"projeto-teste/enum"
	"projeto-teste/model"
	"time"
)

type ProductResponse struct {
	ID            uint                 `json:"id"`
	CategoryID    *int64                `json:"category_id"`
	SubcategoryID *int64                `json:"subcategory_id"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	Name          string               `json:"name"`
	Brand         string               `json:"brand"`
	Description   string               `json:"description"`
	SellingPrice  int64                `json:"selling_price"`
	CostPrice     int64                `json:"cost_price"`
	Stock         int64                `json:"stock"`
	Unit          enum.MeasurementUnit `json:"unit"`
}

func (p *ProductResponse) FromModel(m *model.Product) {
	p.ID = m.ID
	p.CategoryID = m.CategoryID
	p.SubcategoryID = m.SubcategoryID
	p.CreatedAt = m.CreatedAt
	p.UpdatedAt = m.UpdatedAt
	p.Name = m.Name
	p.Brand = m.Brand
	p.Description = m.Description
	p.SellingPrice = m.SellingPrice
	p.CostPrice = m.CostPrice
	p.Stock = m.Stock
	p.Unit = m.Unit
}