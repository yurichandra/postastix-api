package object

import "net/http"

// CategoryResponse represents response of category in json
type CategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Render preprocess before struct rendered
func (res *CategoryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
