package report

type Revenue struct {
	TotalRevenue int `json:"total_revenue"`
	TotalTransaction int `json:"total_transaksi"`
	ProdukTerlaris ProdukTerlaris `json:"produk_terlaris"`
}

type ProdukTerlaris struct {
	Name string `json:"nama"`
	QtySold int   `json:"qty_terjual"`
}