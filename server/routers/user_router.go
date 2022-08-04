package routers

import (
	"go-chi-api/controllers"

	"github.com/go-chi/chi/v5"
)

func UserRouter(router *chi.Mux) {
	router.Get("/api/getExample", controllers.GetUser)
}
