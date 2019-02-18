package object

import (
	"net/http"
)

// UserResponse represents response object for User model.
type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Render preprocesses before struct is rendered.
func (res *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
