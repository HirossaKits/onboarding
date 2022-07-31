package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting Server....")

	router := chi.NewRouter()
	router.Get("/api/getExample", getHandler)
	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("Server is listening on port 8080...")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got me")
}
