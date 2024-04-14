package product

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/google/uuid"
)

type ProductRequest struct {
	Code          string		`validate:"required" form:"code"`
	Name          string		`validate:"required" form:"name"`
	Unit          string		`validate:"required" form:"unit"`
	CategoryId    uuid.UUID	`validate:"required" form:"category_id"`
	Stock         int				`validate:"numeric" form:"stock"`
	PurchasePrice int64			`validate:"numeric" form:"purchase_price"`
	SellingPrice  int64			`validate:"required,numeric" form:"selling_price"`
	Description   string		`validate:"" form:"description"`
	Photo         string		`validate:"" form:"photo"`
}

func ProductRequestToProduct(productRequest ProductRequest) model.Product {
	return model.Product{
		Code: productRequest.Code,
		Name: productRequest.Name,
		Unit: productRequest.Unit,
		CategoryId: productRequest.CategoryId,
		Stock: productRequest.Stock,
		PurchasePrice: productRequest.PurchasePrice,
		SellingPrice: productRequest.SellingPrice,
		Description: productRequest.Description,
		Photo: productRequest.Photo,
	}
}