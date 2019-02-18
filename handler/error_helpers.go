package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	HTTPStatusCode int    `json:"-"`
	Error          string `json:"error"`
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
		HTTPStatusCode: http.StatusNotFound,
		Error:          message,
	}
}

func createBadRequestResponse(message string) *errorResponse {
	if message == "" {
		message = "Bad request"
	}

	return &errorResponse{
		HTTPStatusCode: http.StatusBadRequest,
		Error:          message,
	}
}

func createUnprocessableEntityResponse(message string) *errorResponse {
	if message == "" {
		message = "Unprocessable entity"
	}

	return &errorResponse{
		HTTPStatusCode: http.StatusUnprocessableEntity,
		Error:          message,
	}
}

func createInternalServerErrorResponse(message string) *errorResponse {
	if message == "" {
		message = "Internal server error"
	}

	return &errorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Error:          message,
	}
}
