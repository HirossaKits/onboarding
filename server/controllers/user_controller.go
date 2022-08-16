package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"go-chi-api/db"
	"go-chi-api/ent"
	"go-chi-api/services"
	"go-chi-api/validators"
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

	defer c.client.Close()

	return func(w http.ResponseWriter, r *http.Request) {

		params, err := validators.ValidateGetUserParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			// TODO: エラーメッセージ
			w.Write([]byte(err.Error()))
			return
		}

		user, err := services.GetUserById(ctx, c.client, &params)

		if err != nil {
			w.WriteHeader(500)
			// TODO: エラーメッセージ
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func (c *userControll) CreateUser(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	defer c.client.Close()

	return func(w http.ResponseWriter, r *http.Request) {

		params, err := validators.ValidateCreateUserParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			// TODO: エラーメッセージ
			w.Write([]byte(err.Error()))
			return
		}

		user, err := services.CreateUser(ctx, c.client, &params)

		if err != nil {
			w.WriteHeader(500)
			// TODO: エラーメッセージ
			w.Write([]byte(err.Error()))
			return
		}

		// json.NewEncoder(w).Encode(user)
		json, _ := json.Marshal(user)
		w.Write(json)
	}
}
