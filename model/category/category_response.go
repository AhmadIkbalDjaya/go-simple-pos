package category

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/google/uuid"
)

type CategoryResponse struct {
	ID   uuid.UUID  `json:"id"`
	Name string 		`json:"name"`
}

func ToCategoryResponse(category model.Category) CategoryResponse{
	return CategoryResponse{
		ID: category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []model.Category) []CategoryResponse {
	var categoryResponses = []CategoryResponse{}
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
