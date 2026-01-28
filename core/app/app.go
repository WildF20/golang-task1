package app

import (
	"net/http"

	"golang-task1/middleware"
	"golang-task1/routes"
)

func Run() error {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	server := http.ListenAndServe(":8080", middleware.Logging(mux))

	return server
}