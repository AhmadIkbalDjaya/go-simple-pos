package requests

import "github.com/AhmadIkbalDjaya/go-simple-pos/models"

type CategoryRequest struct {
	Name string `validate:"required" json:"name"`
}

func CategoryRequestToCategory(categoryRequest CategoryRequest) models.Category {
	return models.Category{
		Name: categoryRequest.Name,
	}
}