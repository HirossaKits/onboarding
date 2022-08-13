package controllers

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go-chi-api/db"
	"go-chi-api/ent"
	"go-chi-api/services"
)

type userControll struct {
	client *ent.Client
}

func NewUserController() *userControll {
	return &userControll{
		client: db.NewClient(),
	}
}

type reqBody struct {
	user_id string
}

func (c *userControll) GetUserById(ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {

	defer c.client.Close()

	return func(w http.ResponseWriter, r *http.Request) {

		body := reqBody{}

		bin, err := ioutil.ReadAll(r.Body)
		reader := bytes.NewReader(bin)
		binary.Read(reader, binary.LittleEndian, body)

		user, err := services.GetUserById(ctx, c.client, body)

		if err != nil {

		}
		json.NewEncoder(w).Encode(user)
	}
}
