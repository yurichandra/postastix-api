package main

import (
	"encoding/json"
	"net/http"

	"github.com/dewadg/postastix-api/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func createRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		payload := "Postastix API v1"
		res, _ := json.Marshal(payload)

		w.Write(res)
	})

	// User routes
	router.Route("/v1/users", func(r chi.Router) {
		r.Get("/", handler.GetAllUsers)
	})

	return router
}
