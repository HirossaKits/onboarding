package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

	"go-chi-api/db"
	"go-chi-api/services"
	"go-chi-api/validators"

	"github.com/go-chi/chi/v5"
)

func GetUserById() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

		user_id := chi.URLParam(r, "user_id")
		user, err := services.GetUserById(&ctx, s.Client, user_id)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}

		json, _ := json.Marshal(user)
		w.WriteHeader(500)
		w.Write(json)
	}
}

func CreateUser() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

		// TODO: デバッグ用
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(dump))

		params, err := validators.ValidateCreateUserParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		user, err := services.CreateUser(&ctx, s.Client, &params)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json, _ := json.Marshal(user)
		w.Write(json)
	}
}

func UpdateUser() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

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
		user, err := services.UpdateUser(&ctx, s.Client, user_id, &params)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json, _ := json.Marshal(user)
		w.Write(json)
		w.WriteHeader(204)
	}
}

func DeleteUser() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

		user_id := chi.URLParam(r, "user_id")
		err := services.DeleteUser(&ctx, s.Client, user_id)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(204)
	}
}
