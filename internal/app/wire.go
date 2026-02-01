package app

import (
	"log"
	"database/sql"
	"net/http"

	"golang-task1/internal/category"
	"golang-task1/internal/product"
)

func newWireServer(mux *http.ServeMux, db *sql.DB) error {
	categoryRepo := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepo)
	categoryHandler := category.NewCategoryHandler(categoryService)

	productRepo := product.NewProductRepository(db)
	productService := product.NewProductService(categoryService, productRepo)
	productHandler := product.NewProductHandler(productService)

	category.RegisterRoutes(mux, categoryHandler)
	log.Println("Category module wired successfully")

	product.RegisterRoutes(mux, productHandler)
	log.Println("Product module wired successfully")
	
	return nil
}