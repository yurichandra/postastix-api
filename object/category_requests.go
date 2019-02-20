package object

import (
	"errors"
	"net/http"
)

// CategoryRequest represent request object
type CategoryRequest struct {
	Name string `json:"name"`
}

// Bind for validation of request object in store request
func (req *CategoryRequest) Bind(r *http.Request) error {
	if req.Name == "" {
		return errors.New("Field `name` can't be empty")
	}

	return nil
}
