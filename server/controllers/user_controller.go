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

type userControll struct {
	client *ent.Client
}

func NewUserController() *userControll {
	return &userControll{
		client: db.NewClient(),
	}
}

func (c *userControll) GetUserById(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		defer c.client.Close()

		user_id := chi.URLParam(r, "user_id")
		user, err := services.GetUserById(ctx, c.client, user_id)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func (c *userControll) CreateUser(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		defer c.client.Close()

		// TODO: デバッグ用
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(dump))

		params, err := validators.ValidateCreateUserParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		user, err := services.CreateUser(ctx, c.client, &params)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json, _ := json.Marshal(user)
		w.Write(json)
	}
}

func (c *userControll) UpdateUser(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		defer c.client.Close()

		// TODO: デバッグ用
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(dump))

		params, err := validators.ValidateUpdateUserParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		user_id := chi.URLParam(r, "user_id")
		user, err := services.UpdateUser(ctx, c.client, user_id, &params)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json, _ := json.Marshal(user)
		w.WriteHeader(204)
		w.Write(json)
	}
}

func (c *userControll) DeleteUser(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		defer c.client.Close()

		user_id := chi.URLParam(r, "user_id")
		err := services.DeleteUser(ctx, c.client, user_id)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(204)
	}
}
