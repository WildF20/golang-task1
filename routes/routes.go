package routes

import (
	"net/http"

	"golang-task1/internal/category"
)

func RegisterRoutes(mux *http.ServeMux) {
	category.RegisterCategoryRoutes(mux)
}