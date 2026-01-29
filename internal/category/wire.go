package category

import (
	"net/http"
)

func RegisterCategoryWire(mux *http.ServeMux) {
	RegisterCategoryRoutes(mux)
}