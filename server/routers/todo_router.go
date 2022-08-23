package routers

import (
	"go-chi-api/controllers"

	"github.com/go-chi/chi/v5"
)

type todoRouter struct{}

func TodoRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", controllers.GetTodos())
	r.Get("/{todo_id}", controllers.GetTodoById())
	r.Post("/", controllers.CreateTodo())
	r.Put("/{todo_id}", controllers.UpdateTodo())
	r.Delete("/{todo_id}", controllers.DeleteTodo())
	return r
}
