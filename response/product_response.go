package response

import (
	"projeto-teste/enum"
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

func (p ProductResponse) Print() {
}