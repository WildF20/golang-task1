package category

import (
	"net/http"
)

func RegisterCategoryWire(mux *http.ServeMux) {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)
	RegisterRoutes(mux, handler)
}