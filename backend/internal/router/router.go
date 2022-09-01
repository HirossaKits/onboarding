package router

import (
	"go-chi-api/internal/controller"
	"go-chi-api/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api", func(api chi.Router) {

		api.Route("/user", func(user chi.Router) {
			user.Use(middleware.AuthMiddleWare)
			user.Get("/{user_id}", controller.GetUserById())
			user.Post("/", controller.CreateUser())
			user.Put("/{user_id}", controller.UpdateUser())
			user.Delete("/{user_id}", controller.DeleteUser())
		})
		api.Route("/todo", func(todo chi.Router) {
			todo.Use(middleware.AuthMiddleWare)
			todo.Get("/", controller.GetTodos())
			todo.Get("/{todo_id}", controller.GetTodoById())
			todo.Post("/", controller.CreateTodo())
			todo.Put("/{todo_id}", controller.UpdateTodo())
			todo.Delete("/{todo_id}", controller.DeleteTodo())
		})
	})
	return r
}

// func todoRoute() *chi.Mux {
// 	r := chi.NewRouter()
// 	r.Use(middleware.AuthMiddleWare)
// 	r.Get("/", controller.GetTodos())
// 	r.Get("/{todo_id}", controller.GetTodoById())
// 	r.Post("/", controller.CreateTodo())
// 	r.Put("/{todo_id}", controller.UpdateTodo())
// 	r.Delete("/{todo_id}", controller.DeleteTodo())
// 	return r
// }

// func userRoute() *chi.Mux {
// 	r := chi.NewRouter()
// 	r.Use(middleware.AuthMiddleWare)
// 	r.Get("/{user_id}", controller.GetUserById())
// 	r.Post("/", controller.CreateUser())
// 	r.Put("/{user_id}", controller.UpdateUser())
// 	r.Delete("/{user_id}", controller.DeleteUser())
// 	return r
// }
