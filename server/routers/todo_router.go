package routers

import (
	"go-chi-api/controllers"
	"go-chi-api/middlewares"

	"github.com/go-chi/chi/v5"
)

type todoRouter struct{}

func TodoRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.AuthMiddleWare)
	r.Get("/", controllers.GetTodos())
	r.Get("/{todo_id}", controllers.GetTodoById())
	r.Post("/", controllers.CreateTodo())
	r.Put("/{todo_id}", controllers.UpdateTodo())
	r.Delete("/{todo_id}", controllers.DeleteTodo())
	return r
}
