package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go-chi-api/db"
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

	s.router.Route("/api", func(r chi.Router) {
		userRouter := routers.NewUserRouter()
		r.Mount("/user", userRouter.Route(&ctx))
	})
}

func main() {

	args := os.Args

	if len(args) == 2 {
		if args[1] == "migrate" {
			db.Migrate()
		}
	}

	if len(args) == 1 {

		s := New()
		s.route()

		err := http.ListenAndServe(":8080", s.router)
		if err != nil {
			log.Fatalf("failed starting server: %v", err)
		}
	}
}
