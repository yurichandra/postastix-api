package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type app struct {
	db     *gorm.DB
	router *chi.Mux
}

func createApp(db *gorm.DB, router *chi.Mux) *app {
	return &app{
		db:     db,
		router: router,
	}
}

func (a *app) Run() {
	fmt.Printf("App running on port %s\n", os.Getenv("APP_PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), a.router)
}
