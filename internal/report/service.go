package report

import (
	"time"
	"context"
)

type ReportService struct {
	repo *ReportRepository
}

func NewReportService(repo *ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetRevenue(ctx context.Context, startDate, endDate time.Time) (Revenue, error) {
	return s.repo.FetchRevenue(ctx, startDate, endDate)
}