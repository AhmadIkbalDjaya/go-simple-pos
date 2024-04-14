package product

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/category"
	"github.com/google/uuid"
)

type ProductResponse struct {
	ID            uuid.UUID				`json:"id"`
	Code          string					`json:"code"`
	Name          string					`json:"name"`
	Unit          string					`json:"unit"`
	Category      category.CategoryResponse	`json:"category"`
	Stock         int							`json:"stock"`
	PurchasePrice int64						`json:"purchase_price"`
	SellingPrice  int64						`json:"selling_price"`
	Description   string					`json:"description"`
	Photo         string					`json:"photo"`
}

func ToProductResponse(product model.Product) ProductResponse {
	return ProductResponse{
		ID: product.ID,
		Code: product.Code,
		Name: product.Name,
		Unit: product.Unit,
		Category: category.ToCategoryResponse(product.Category),
		Stock: product.Stock,
		PurchasePrice: product.PurchasePrice,
		SellingPrice: product.SellingPrice,
		Description: product.Description,
		Photo: product.Photo,
	}
}

func ToProductResponses(products []model.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}