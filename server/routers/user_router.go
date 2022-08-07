package routers

import (
	"go-chi-api/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
}

func UserRouter(*ctx) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{user_id}", controllers.GetUser)

	return r
}
