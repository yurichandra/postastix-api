package main

import (
	"postastix-api/handler"

	"github.com/go-chi/chi"
)

func createRouter() chi.Router {
	r := chi.NewRouter()

	r.Mount("/users", handler.UserRoutes())

	return r
}
