package app

import (
	"net/http"

	"golang-task1/config"
	"golang-task1/internal/shared/middleware"
)

func newHttpServer(mux *http.ServeMux) error {
	cfg, _ := config.LoadConfig()

	handler := middleware.Chain(
		mux,
		middleware.Logging,
		middleware.APIKey(cfg.APIKey),
	)

	err := http.ListenAndServe(cfg.Address, handler)

	return err
}