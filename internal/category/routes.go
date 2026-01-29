package category

import (
	"net/http"
)

func RegisterCategoryRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /categories", GetAll)
	mux.HandleFunc("GET /categories/{id}", GetByID)
	mux.HandleFunc("POST /categories", Create)
	mux.HandleFunc("PUT /categories/{id}", Update)
	mux.HandleFunc("DELETE /categories/{id}", Delete)
}