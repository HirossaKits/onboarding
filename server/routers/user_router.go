package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(router *chi.Mux) {
	router.Get("/api/getExample", getHandler)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
