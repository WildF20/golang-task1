package category

import (
	"log"
	"net/http"
	"database/sql"
)

func RegisterCategoryWire(mux *http.ServeMux, db *sql.DB) {
	repo := NewCategoryRepository(db)
	service := NewCategoryService(repo)
	handler := NewCategoryHandler(service)
	RegisterRoutes(mux, handler)

	log.Println("Product module wired successfully")
}