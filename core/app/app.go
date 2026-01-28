package app

import (
	"net/http"

	"golang-task1/config"
	"golang-task1/middleware"
	"golang-task1/routes"
)

func Run() error {
	cfg, _ := config.LoadConfig()

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	server := http.ListenAndServe(cfg.Address, middleware.Logging(mux))

	return server
}