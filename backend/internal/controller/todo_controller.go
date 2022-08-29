package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

	"go-chi-api/internal/db"
	"go-chi-api/internal/service"
	"go-chi-api/internal/validator"

	"github.com/go-chi/chi/v5"
)

func GetTodos() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

		// TODO: デバッグ用
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(dump))

		params, err := validator.ValidateGetTodosParam(&r.Body)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(400)
			return
		}
		Todo, err := service.GetTodos(&ctx, s.Client, &params)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(Todo)
	}
}

func GetTodoById() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

		todo_id := chi.URLParam(r, "todo_id")
		todo, err := service.GetTodoById(&ctx, s.Client, todo_id)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json, _ := json.Marshal(todo)
		w.Write(json)
	}
}

func CreateTodo() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

		// TODO: デバッグ用
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(dump))

		params, err := validator.ValidateCreateTodoParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		Todo, err := service.CreateTodo(&ctx, s.Client, &params)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json, err := json.Marshal(Todo)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(json)
	}
}

func UpdateTodo() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

		// TODO: デバッグ用
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(dump))

		params, err := validator.ValidateUpdateTodoParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		todo_id := chi.URLParam(r, "todo_id")
		todo, err := service.UpdateTodo(&ctx, s.Client, todo_id, &params)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json, _ := json.Marshal(todo)
		w.Write(json)
		w.WriteHeader(204)
	}
}

func DeleteTodo() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

		todo_id := chi.URLParam(r, "todo_id")
		err := service.DeleteTodo(&ctx, s.Client, todo_id)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(204)
	}
}
