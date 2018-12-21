package handler

import (
	"fmt"
	"net/http"

	"github.com/dewadg/postastix-api/db"
	"github.com/go-chi/render"
)

type userResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type userListResponse []*userListResponse

func (u *userResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func createUserReponse(user *db.User) *userResponse {
	return &userResponse{
		ID:        user.ID,
		Name:      user.Name,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func createUserListResponse(users []*db.User) []render.Renderer {
	list := make([]render.Renderer, 0)

	for _, user := range users {
		list = append(list, createUserReponse(user))
	}

	return list
}

// GetAllUsers retrieves users and displays it as JSON.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	payload := createUserListResponse(userService.Get())

	if err := render.RenderList(w, r, payload); err != nil {
		fmt.Println("Error")
		return
	}
}
