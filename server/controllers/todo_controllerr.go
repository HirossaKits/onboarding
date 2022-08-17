package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

	"go-chi-api/db"
	"go-chi-api/ent"
	"go-chi-api/services"
	"go-chi-api/validators"

	"github.com/go-chi/chi/v5"
)

type todoController struct {
	client *ent.Client
}

func NewTodoController() *todoController {
	return &todoController{
		client: db.NewClient(),
	}
}

// func (c *todoController) GetTodos(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		defer r.Body.Close()
// 		defer c.client.Close()

// 		todo_id := chi.URLParam(r, "todo_id")
// 		Todo, err := services.GetTodos(ctx, c.client)

// 		if err != nil {
// 			w.WriteHeader(500)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}

// 		json.NewEncoder(w).Encode(Todo)
// 	}
// }

func (c *todoController) GetTodoById(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		defer c.client.Close()

		todo_id := chi.URLParam(r, "todo_id")
		todo, err := services.GetTodoById(ctx, c.client, todo_id)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(todo)
	}
}

func (c *todoController) CreateTodo(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		defer c.client.Close()

		// TODO: デバッグ用
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(dump))

		params, err := validators.ValidateCreateTodoParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		Todo, err := services.CreateTodo(ctx, c.client, &params)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json, _ := json.Marshal(Todo)
		w.Write(json)
	}
}

func (c *todoController) UpdateTodo(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		defer c.client.Close()

		// TODO: デバッグ用
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(dump))

		params, err := validators.ValidateUpdateTodoParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		todo_id := chi.URLParam(r, "todo_id")
		todo, err := services.UpdateTodo(ctx, c.client, todo_id, &params)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json, _ := json.Marshal(todo)
		w.WriteHeader(204)
		w.Write(json)
	}
}

func (c *todoController) DeleteTodo(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		defer c.client.Close()

		todo_id := chi.URLParam(r, "todo_id")
		err := services.DeleteTodo(ctx, c.client, todo_id)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(204)
	}
}
