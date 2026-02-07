package report

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *ReportHandler) {
	mux.HandleFunc("GET /api/report", handler.GetRevenue)
	mux.HandleFunc("POST /api/report/hari-ini", handler.GetTodayRevenue)
}