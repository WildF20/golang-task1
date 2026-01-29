package app

import (
	"net/http"

	"golang-task1/config"
	"golang-task1/internal/shared/middleware"
)

func newHttpServer() error {
	cfg, _ := config.LoadConfig()
	mux := http.NewServeMux()

	if errWire := newWireServer(mux); errWire != nil {
		return errWire
	}

	err := http.ListenAndServe(cfg.Address, middleware.Logging(mux))

	return err
}