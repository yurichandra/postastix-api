package handler

import (
	"fmt"
	"net/http"
	"postastix-api/object"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// UserRoutes returns router of user handlers.
func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getUsers)

	return r
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	payload := object.CreateUserListResponse(userService.Get())

	if err := render.RenderList(w, r, payload); err != nil {
		fmt.Println(err.Error())
		return
	}
}
