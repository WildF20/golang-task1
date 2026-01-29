package app

import (
	"net/http"

	"golang-task1/config"
	"golang-task1/internal/shared/middleware"

	"golang-task1/internal/category"
)

func newHttpServer() error {
	cfg, _ := config.LoadConfig()
	mux := http.NewServeMux()

	category.RegisterCategoryWire(mux)

	err := http.ListenAndServe(cfg.Address, middleware.Logging(mux))

	return err
}