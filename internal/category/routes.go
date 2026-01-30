package category

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *CategoryHandler) {
	mux.HandleFunc("GET /categories", handler.GetAll)
	mux.HandleFunc("GET /categories/{id}", handler.GetByID)
	mux.HandleFunc("POST /categories", handler.Create)
	mux.HandleFunc("PUT /categories/{id}", handler.Update)
	mux.HandleFunc("DELETE /categories/{id}", handler.Delete)
}