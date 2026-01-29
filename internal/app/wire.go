package app

import (
	"net/http"

	"golang-task1/internal/category"
)

func newWireServer(mux *http.ServeMux) error {
	category.RegisterCategoryWire(mux)
	
	return nil
}