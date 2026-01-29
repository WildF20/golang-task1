package app

import (
	"log"
)

func Run() error {
	if err := newDBConnection(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	if err := newHttpServer(); err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}

	return nil
}