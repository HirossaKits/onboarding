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

func GetUserById() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		s := db.GetInstance()
		defer r.Body.Close()

		user_id := chi.URLParam(r, "user_id")
		user, err := service.GetUserById(&ctx, s.Client, user_id)

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

		params, err := validator.ValidateCreateUserParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		user, err := service.CreateUser(&ctx, s.Client, &params)

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

		params, err := validator.ValidateUpdateUserParam(&r.Body)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		user_id := chi.URLParam(r, "user_id")
		user, err := service.UpdateUser(&ctx, s.Client, user_id, &params)

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
		err := service.DeleteUser(&ctx, s.Client, user_id)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(204)
	}
}
