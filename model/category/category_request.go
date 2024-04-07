package category

import "github.com/AhmadIkbalDjaya/go-simple-pos/model"

type CategoryRequest struct {
	Name string `validate:"required" json:"name"`
}

func CategoryRequestToCategory(categoryRequest CategoryRequest) model.Category {
	return model.Category{
		Name: categoryRequest.Name,
	}
}