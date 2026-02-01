package app

import (
	"database/sql"
	"net/http"

	"golang-task1/internal/category"
	"golang-task1/internal/product"
)

func newWireServer(mux *http.ServeMux, db *sql.DB) error {
	category.RegisterCategoryWire(mux, db)
	product.RegisterProductWire(mux, db)
	
	return nil
}