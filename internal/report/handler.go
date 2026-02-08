package report

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ReportHandler struct {
	service *ReportService
}

func NewReportHandler(service *ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	layout := "2006-01-02"

	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		http.Error(w, "invalid start_date format (use yyyy-mm-dd)", http.StatusBadRequest)
		log.Println("Error parsing start date:", err)
		return
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		http.Error(w, "invalid end_date format (use yyyy-mm-dd)", http.StatusBadRequest)
		log.Println("Error parsing end date:", err)
		return
	}
	
	revenue, err := h.service.GetRevenue(ctx, startDate, endDate)
	if err != nil {
		http.Error(w, "Failed to get revenue for the given date range", http.StatusInternalServerError)
		log.Println("Error fetching revenue:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(revenue)
}

func (h *ReportHandler) GetTodayRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	today := time.Date(2026, 2, 8, 0, 0, 0, 0, loc)
	tomorrow := time.Now().AddDate(0, 0, 1)

	revenue, err := h.service.GetRevenue(ctx, today, tomorrow)
	if err != nil {
		http.Error(w, "Failed to get today's revenue", http.StatusInternalServerError)
		log.Println("Error fetching today's revenue:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(revenue)
}