package app

import (
	"log"
	"net/http"

	"golang-task1/database"
	"golang-task1/config"
	"golang-task1/middleware"
	"golang-task1/routes"
)

func Run() error {
	cfg, _ := config.LoadConfig()

	db, err := database.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	server := http.ListenAndServe(cfg.Address, middleware.Logging(mux))

	return server
}