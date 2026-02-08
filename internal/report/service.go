package report

type ReportService struct {
	repo *ReportRepository
}

func NewReportService(repo *ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetRevenue(startDate string, endDate string) (Revenue, error) {
	return s.repo.FetchRevenue(startDate, endDate)
}