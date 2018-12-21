package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/dewadg/postastix-api/db"
	"github.com/go-chi/render"
)

type key string

const (
	userInCtx key = "user"
)

type userResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (u *userResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type userListResponse []*userListResponse

type storeUserRequest struct {
	Name            string `json:"name"`
	FullName        string `json:"fullName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (s *storeUserRequest) Bind(r *http.Request) error {
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

// UserCtx fetches user and set it as context value.
func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var idParam string
		if idParam = chi.URLParam(r, "id"); idParam == "" {
			render.Render(w, r, createBadRequestResponse(""))
		}
		id, _ := strconv.ParseUint(idParam, 10, 8)

		user, err := userService.Find(uint(id))
		if err != nil {
			if err.Error() == "User not found" {
				render.Render(w, r, createNotFoundResponse(err.Error()))
				return
			}
		}

		ctx := context.WithValue(r.Context(), userInCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetAllUsers retrieves users and displays it as JSON.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	payload := createUserListResponse(userService.Get())

	if err := render.RenderList(w, r, payload); err != nil {
		fmt.Println("Error")
		return
	}
}

// StoreUser stores a user and displays it as JSON.
func StoreUser(w http.ResponseWriter, r *http.Request) {
	payload := new(storeUserRequest)
	if err := render.Bind(r, payload); err != nil {
		render.Render(w, r, createBadRequestResponse(""))
		return
	}

	if payload.Password != payload.ConfirmPassword {
		render.Render(w, r, createUnprocessableEntityResponse("Password not match"))
		return
	}

	user, err := userService.Create(payload.Name, payload.FullName, payload.Password)
	if err != nil {
		render.Render(w, r, createUnprocessableEntityResponse(err.Error()))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, createUserReponse(user))
}

// GetOneUser retrieves a user and displays it as JSON.
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userInCtx).(*db.User)
	render.Render(w, r, createUserReponse(user))
}
