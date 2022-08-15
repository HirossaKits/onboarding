package routers

import (
	"context"
	"go-chi-api/controllers"

	"github.com/go-chi/chi/v5"
)

type userRouter struct{}

func NewUserRouter() *userRouter {
	return &userRouter{}
}

func (*userRouter) Route(ctx *context.Context) *chi.Mux {
	r := chi.NewRouter()
	c := controllers.NewUserController()
	r.Get("/{user_id}", c.GetUserById(ctx))
	r.Post("/", c.CreateUser(ctx))
	return r
}
