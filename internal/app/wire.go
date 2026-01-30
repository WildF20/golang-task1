package app

import (
	"net/http"
	"database/sql"

	"golang-task1/internal/category"
)

func newWireServer(mux *http.ServeMux, db *sql.DB) error {
	category.RegisterCategoryWire(mux, db)
	
	return nil
}