package routers

import (
	"context"
	"go-chi-api/controllers"

	"github.com/go-chi/chi/v5"
)

type todoRouter struct{}

func TodoRoute(ctx *context.Context) *chi.Mux {
	r := chi.NewRouter()
	c := controllers.NewTodoController()
	// r.Get("/", c.GetTodos(ctx))
	r.Get("/{todo_id}", c.GetTodoById(ctx))
	r.Post("/", c.CreateTodo(ctx))
	r.Put("/{todo_id}", c.UpdateTodo(ctx))
	r.Delete("/{todo_id}", c.DeleteTodo(ctx))
	return r
}
