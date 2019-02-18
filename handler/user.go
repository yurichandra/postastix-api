package handler

import (
	"context"
	"fmt"
	"net/http"
	"postastix-api/object"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// UserRoutes returns router of user handlers.
func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getUsers)

	return r
}

func userContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
		user, err := userService.Find(uint(userID))
		if err != nil {
			render.Render(w, r, createNotFoundResponse(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	payload := object.CreateUserListResponse(userService.Get())

	if err := render.RenderList(w, r, payload); err != nil {
		fmt.Println(err.Error())
		return
	}
}
