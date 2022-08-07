package main

import (
	"go-chi-api/ent"
	"log"
	"net/http"

	"go-chi-api/db"
	"go-chi-api/routers"

	"github.com/go-chi/chi/v5"
)

type server struct {
	client *ent.Client
	router *chi.Mux
}

func New() *server {
	return &server{
		client: db.NewClient(),
		router: chi.NewRouter(),
	}
}

ctx := context.Background()
ctx.server := server

func (s *server) Route() {
	s.router.Route("/api", func(r chi.Router) {
		r.Route("/user", routers.UserRouter())
		r.Route("/todo", routers.UserRouter())
	})
}

func main() {

	s := New()
	s.Route()

	defer s.client.Close()

	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		log.Fatalf("failed starting server: %v", err)
	}
}
