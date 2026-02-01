package product

import (
	"log"
	"net/http"
	"database/sql"
)

func RegisterProductWire(mux *http.ServeMux, db *sql.DB) {
	repo := NewProductRepository(db)
	service := NewProductService(repo)
	handler := NewProductHandler(service)
	RegisterRoutes(mux, handler)

	log.Println("Product module wired successfully")
}