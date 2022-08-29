package main

import (
	"log"
	"net/http"

	"go-chi-api/internal/router"
)

func main() {
	err := http.ListenAndServe(":8080", router.Route())
	if err != nil {
		log.Fatalf("failed starting server: %v", err)
	}
}
