package main

import (
	"log"

	"golang-task1/core/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
