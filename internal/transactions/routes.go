package transactions

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *TransactionHandler) {
	mux.HandleFunc("POST /api/checkout", handler.Checkout)
}