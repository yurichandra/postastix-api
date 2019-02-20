package object

import (
	"postastix-api/model"

	"github.com/go-chi/render"
)

// CreateCategoryResponse is creating response for category model
func CreateCategoryResponse(payload model.Category) render.Renderer {
	return &CategoryResponse{
		ID:        payload.ID,
		Name:      payload.Name,
		CreatedAt: payload.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: payload.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

//CreateCategoryListResponse is creating response for list of category model
func CreateCategoryListResponse(payload []model.Category) []render.Renderer {
	response := make([]render.Renderer, 0)

	for _, item := range payload {
		response = append(response, CreateCategoryResponse(item))
	}

	return response
}
