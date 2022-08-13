package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go-chi-api/routers"

	"github.com/go-chi/chi/v5"
)

type server struct {
	router *chi.Mux
}

func New() *server {
	return &server{
		router: chi.NewRouter(),
	}
}

func (s *server) route() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	s.router.Route("/api", func(r chi.Router) {
		userRouter := routers.NewUserRouter()
		r.Mount("/user", userRouter.Route(&ctx))
	})
}

func main() {

	s := New()
	s.route()

	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		log.Fatalf("failed starting server: %v", err)
	}
}
