package handler

import (
	"context"
	"fmt"
	"net/http"
	"postastix-api/model"
	"postastix-api/object"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// UserRoutes returns router of user handlers.
func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getUsers)
	r.Route("/{userID}", func(r chi.Router) {
		r.Use(userContext)
		r.Get("/", getUser)
	})

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

func getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(userCtx).(model.User)
	if !ok {
		render.Render(w, r, createUnprocessableEntityResponse(""))
		return
	}

	render.Render(w, r, object.CreateUserResponse(user))
}
