package report

import (
	"database/sql"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) FetchRevenue(startDate string, endDate string) (Revenue, error) {
	return Revenue{}, nil
}