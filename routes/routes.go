package routes

import (
	"net/http"

	"golang-task1/handler"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /categories", handler.GetAllCategories)
	mux.HandleFunc("GET /categories/{id}", handler.GetCategoryByID)
	mux.HandleFunc("POST /categories", handler.CreateCategory)
	mux.HandleFunc("PUT /categories/{id}", handler.UpdateCategory)
	mux.HandleFunc("DELETE /categories/{id}", handler.DeleteCategory)
}