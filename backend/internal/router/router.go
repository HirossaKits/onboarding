package router

import (
	"go-chi-api/internal/controller"
	"go-chi-api/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func Route() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Mount("/user", UserRoute())
		r.Mount("/todo", TodoRoute())
	})
	return r
}

func TodoRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleWare)
	r.Get("/", controller.GetTodos())
	r.Get("/{todo_id}", controller.GetTodoById())
	r.Post("/", controller.CreateTodo())
	r.Put("/{todo_id}", controller.UpdateTodo())
	r.Delete("/{todo_id}", controller.DeleteTodo())
	return r
}

func UserRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleWare)
	r.Get("/{user_id}", controller.GetUserById())
	r.Post("/", controller.CreateUser())
	r.Put("/{user_id}", controller.UpdateUser())
	r.Delete("/{user_id}", controller.DeleteUser())
	return r
}
