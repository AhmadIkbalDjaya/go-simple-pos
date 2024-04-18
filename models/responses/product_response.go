package responses

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/models"
	"github.com/google/uuid"
)

type ProductResponse struct {
	ID            uuid.UUID				`json:"id"`
	Code          string					`json:"code"`
	Name          string					`json:"name"`
	Unit          string					`json:"unit"`
	Category      CategoryResponse	`json:"category"`
	Stock         int							`json:"stock"`
	PurchasePrice int64						`json:"purchase_price"`
	SellingPrice  int64						`json:"selling_price"`
	Description   string					`json:"description"`
	Photo         string					`json:"photo"`
}

func ToProductResponse(product models.Product) ProductResponse {
	return ProductResponse{
		ID: product.ID,
		Code: product.Code,
		Name: product.Name,
		Unit: product.Unit,
		Category: ToCategoryResponse(product.Category),
		Stock: product.Stock,
		PurchasePrice: product.PurchasePrice,
		SellingPrice: product.SellingPrice,
		Description: product.Description,
		Photo: product.Photo,
	}
}

func ToProductResponses(products []models.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}