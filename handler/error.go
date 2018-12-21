package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	HTTPStatusCode int    `json:"-"`
	Error          string `json:"errors"`
}

func (e *errorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func createNotFoundResponse(message string) *errorResponse {
	if message == "" {
		message = "Resource not found"
	}

	return &errorResponse{
		HTTPStatusCode: 404,
		Error:          message,
	}
}

func createBadRequestResponse(message string) *errorResponse {
	if message == "" {
		message = "Bad request"
	}

	return &errorResponse{
		HTTPStatusCode: 401,
		Error:          message,
	}
}
