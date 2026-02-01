package product

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *ProductHandler) {
	mux.HandleFunc("GET /products", handler.GetAll)
	mux.HandleFunc("GET /products/{id}", handler.GetByID)
	mux.HandleFunc("POST /products", handler.Create)
	mux.HandleFunc("PUT /products/{id}", handler.Update)
	mux.HandleFunc("DELETE /products/{id}", handler.Delete)
}