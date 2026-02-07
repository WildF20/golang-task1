package app

import (
	"database/sql"
	"log"
	"net/http"

	"golang-task1/internal/category"
	"golang-task1/internal/product"
	"golang-task1/internal/report"
	"golang-task1/internal/transactions"
)

func newWireServer(mux *http.ServeMux, db *sql.DB) error {
	categoryRepo := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepo)
	categoryHandler := category.NewCategoryHandler(categoryService)

	productRepo := product.NewProductRepository(db)
	productService := product.NewProductService(categoryService, productRepo)
	productHandler := product.NewProductHandler(productService)

	transactionRepo := transactions.NewTransactionRepository(db)
	transactionService := transactions.NewTransactionService(transactionRepo)
	transactionHandler := transactions.NewTransactionHandler(transactionService)

	reportRepo := report.NewReportRepository(db)
	reportService := report.NewReportService(reportRepo)
	reportHandler := report.NewReportHandler(reportService)
	
	category.RegisterRoutes(mux, categoryHandler)
	log.Println("Category module wired successfully")

	product.RegisterRoutes(mux, productHandler)
	log.Println("Product module wired successfully")

	transactions.RegisterRoutes(mux, transactionHandler)
	log.Println("Transaction module wired successfully")

	report.RegisterRoutes(mux, reportHandler)
	log.Println("Report module wired successfully")
	
	return nil
}