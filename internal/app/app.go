package app

import (
	"log"

	"net/http"
)

func Run() error {
	db, err := newDBConnection();
	if  err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	mux := http.NewServeMux()
	if err := newWireServer(mux, db); err != nil {
		log.Fatal("Failed to wire services:", err)
	}

	if err := newHttpServer(mux); err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}

	return nil
}