package report

import (
	"database/sql"
	"time"
	"context"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) FetchRevenue(ctx context.Context, startDate, endDate time.Time) (Revenue, error) {
	var revenue Revenue
	var produkTerlaris ProdukTerlaris

	err := r.db.QueryRowContext(ctx, `
		select 
			p.name as nama,
			count(*) as qty_terjual
		from transaction_details as td
		inner join transactions as t on td.transaction_id = t.id
		inner join products as p on 
			td.product_id = p.id
		where
			t.created_at >= $1
			and t.created_at < $2
		group by
			p.name
		order by
			qty_terjual desc
		limit 1
	`, startDate, endDate).Scan(&produkTerlaris.Name, &produkTerlaris.QtySold)

	if err != nil {
		return revenue, err
	}

	err = r.db.QueryRowContext(ctx, `
		select 
			sum(t.total_amount) as total_revenue,
			count(*) as total_transaksi
		from 
			transactions as t
		where
			t.created_at >= $1
			and t.created_at < $2
	`, startDate, endDate).Scan(&revenue.TotalRevenue, &revenue.TotalTransaction)

	revenue.ProdukTerlaris = produkTerlaris

	if err != nil {
		return revenue, err
	}

	return revenue, nil
}