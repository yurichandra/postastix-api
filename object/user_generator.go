package object

import (
	"postastix-api/model"

	"github.com/go-chi/render"
)

// CreateUserResponse returns UserResponse from User model.
func CreateUserResponse(payload model.User) render.Renderer {
	return &UserResponse{
		ID:        payload.ID,
		Name:      payload.Name,
		FullName:  payload.FullName,
		CreatedAt: payload.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: payload.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// CreateUserListResponse returns list of UserResponse.
func CreateUserListResponse(payload []model.User) []render.Renderer {
	response := make([]render.Renderer, 0)

	for _, item := range payload {
		response = append(response, CreateUserResponse(item))
	}

	return response
}
