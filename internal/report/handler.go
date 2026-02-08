package report

import (
	"net/http"
	"time"
	"encoding/json"
)

type ReportHandler struct {
	service *ReportService
}

func NewReportHandler(service *ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetRevenue(w http.ResponseWriter, r *http.Request) {
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	layout := "2006-01-02"

	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		http.Error(w, "invalid start_date format (use yyyy-mm-dd)", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		http.Error(w, "invalid end_date format (use yyyy-mm-dd)", http.StatusBadRequest)
		return
	}
	
	revenue, err := h.service.GetRevenue(startDate.Format(layout), endDate.Format(layout))
	if err != nil {
		http.Error(w, "Failed to get revenue for the given date range", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(revenue)
}

func (h *ReportHandler) GetTodayRevenue(w http.ResponseWriter, r *http.Request) {
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	revenue, err := h.service.GetRevenue(today, tomorrow)
	if err != nil {
		http.Error(w, "Failed to get today's revenue", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(revenue)
}