package product

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/google/uuid"
)

type ProductRequest struct {
	Code          string		`validate:"required" json:"code"`
	Name          string		`validate:"required" json:"name"`
	Unit          string		`validate:"required" json:"unit"`
	CategoryId    uuid.UUID	`validate:"required" json:"category"`
	Stock         int				`validate:"numeric" json:"stock"`
	PurchasePrice int64			`validate:"numeric" json:"purchase_price"`
	SellingPrice  int64			`validate:"required,numeric" json:"selling_price"`
	Description   string		`validate:"" json:"description"`
	Photo         string		`validate:"" json:"photo"`
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