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

// CategoryRoutes return router for Category handler
func CategoryRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getCategories)
	r.Post("/", storeCategory)
	r.Route("/{categoryID}", func(r chi.Router) {
		r.Use(categoryContext)
		r.Get("/", findCategory)
		r.Patch("/", updateCategory)
		r.Delete("/", destroyCategory)
	})

	return r
}

func categoryContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoryID, _ := strconv.Atoi(chi.URLParam(r, "categoryID"))
		category, err := categoryService.Find(uint(categoryID))

		if err != nil {
			render.Render(w, r, createNotFoundResponse(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), categoryCtx, category)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	payload := object.CreateCategoryListResponse(categoryService.Get())

	if err := render.RenderList(w, r, payload); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func storeCategory(w http.ResponseWriter, r *http.Request) {
	payload := object.CategoryRequest{}

	if err := render.Bind(r, &payload); err != nil {
		createUnprocessableEntityResponse(err.Error())
		return
	}

	category, err := categoryService.Create(payload.Name)

	if err != nil {
		createUnprocessableEntityResponse(err.Error())
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, object.CreateCategoryResponse(category))
}

func findCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value(categoryCtx).(model.Category)

	if !ok {
		createUnprocessableEntityResponse("")
		return
	}

	render.Render(w, r, object.CreateCategoryResponse(category))
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value(categoryCtx).(model.Category)
	if !ok {
		createUnprocessableEntityResponse("")
		return
	}

	payload := object.CategoryRequest{}
	if err := render.Bind(r, &payload); err != nil {
		createUnprocessableEntityResponse(err.Error())
		return
	}

	cat, err := categoryService.Update(category.ID, payload.Name)
	if err != nil {
		render.Render(w, r, createUnprocessableEntityResponse(err.Error()))
		return
	}

	render.Render(w, r, object.CreateCategoryResponse(cat))
}

func destroyCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value(categoryCtx).(model.Category)

	if !ok {
		createUnprocessableEntityResponse("")
		return
	}

	categoryService.Delete(category.ID)
}
