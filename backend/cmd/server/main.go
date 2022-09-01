package main

import (
	"log"
	"net/http"

	"go-chi-api/internal/router"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	return &Server{
		router: router.NewRouter(),
	}
}

func main() {
	s := NewServer()
	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		log.Fatalf("failed starting server: %v", err)
	}
}
