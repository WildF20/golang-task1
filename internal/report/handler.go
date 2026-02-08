package report

import (
	"net/http"
)

type ReportHandler struct {
	service *ReportService
}

func NewReportHandler(service *ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetRevenue(w http.ResponseWriter, r *http.Request) {
	
}

func (h *ReportHandler) GetTodayRevenue(w http.ResponseWriter, r *http.Request) {

}