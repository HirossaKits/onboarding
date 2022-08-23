package routers

import (
	"go-chi-api/controllers"

	"github.com/go-chi/chi/v5"
)

func UserRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{user_id}", controllers.GetUserById())
	r.Post("/", controllers.CreateUser())
	r.Put("/{user_id}", controllers.UpdateUser())
	r.Delete("/{user_id}", controllers.DeleteUser())
	return r
}
