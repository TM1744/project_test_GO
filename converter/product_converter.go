package converter

// import (
// 	"projeto-teste/model"
// 	"projeto-teste/request"
// 	"projeto-teste/response"
// )

// func ProductRequestToSchema(r *request.ProductRequest) model.Product {
// 	return model.Product{
// 		CategoryID:    r.CategoryID,
// 		SubcategoryID: r.SubcategoryID,
// 		Name:          &r.Name,
// 		Brand:         &r.Brand,
// 		Description:   &r.Description,
// 		SellingPrice:  &r.SellingPrice,
// 		CostPrice:     &r.CostPrice,
// 		Stock:         &r.Stock,
// 		Unit:          &r.Unit,
// 	}
// }

// func ProductSchemaToResponse(s *model.Product) response.ProductResponse {
// 	return response.ProductResponse{
// 		ID:            &s.ID,
// 		CreatedAt:     &s.CreatedAt,
// 		UpdatedAt:     &s.UpdatedAt,
// 		CategoryID:    s.CategoryID,
// 		SubcategoryID: s.SubcategoryID,
// 		Name:          s.Name,
// 		Brand:         s.Brand,
// 		Description:   s.Description,
// 		SellingPrice:  s.SellingPrice,
// 		CostPrice:     s.CostPrice,
// 		Stock:         s.Stock,
// 		Unit:          s.Unit,
// 	}
// }

// func ProductSchemaToSimpleResponse(s *model.Product) response.ProductSimpleResponse {
// 	return response.ProductSimpleResponse{
// 		ID:            &s.ID,
// 		CategoryID:    s.CategoryID,
// 		SubcategoryID: s.SubcategoryID,
// 		Name:          s.Name,
// 		Brand:         s.Brand,
// 		SellingPrice:  s.SellingPrice,
// 		Stock:         s.Stock,
// 	}
// }
