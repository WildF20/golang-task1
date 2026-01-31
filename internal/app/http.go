package app

import (
	"net/http"

	"golang-task1/config"
	"golang-task1/internal/shared/middleware"
)

func newHttpServer(mux *http.ServeMux) error {
	cfg, _ := config.LoadConfig()

	err := http.ListenAndServe(cfg.Address, middleware.Logging(mux))

	return err
}