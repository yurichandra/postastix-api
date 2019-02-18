package main

import (
	"fmt"
	"net/http"
	"os"
	"postastix-api/handler"

	"github.com/go-chi/chi"
)

func createRouter() chi.Router {
	handler.InitServices()
	r := chi.NewRouter()

	r.Mount("/users", handler.UserRoutes())

	return r
}

func serveHTTP() {
	router := createRouter()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("App running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
