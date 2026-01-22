package main

import (
	"log"
	"net/http"

	"golang-task1/middleware"
	"golang-task1/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	if err := http.ListenAndServe(":8080", middleware.Logging(mux)); err != nil {
		log.Fatal(err)
	}
}
