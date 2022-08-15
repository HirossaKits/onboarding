package services

import (
	"context"
	"go-chi-api/ent"
	"go-chi-api/ent/user"
	"go-chi-api/validators"

	"github.com/google/uuid"
)

func GetUserById(ctx *context.Context, client *ent.Client, params *validators.GetUserParams) (string, error) {

	user_id, _ := uuid.Parse(params.User_id)

	user, err := client.User.
		Query().
		Where(user.ID(user_id)).
		Select(user.FieldID, user.FieldName).
		String(*ctx)

	return user, err
}

func CreateUser(ctx *context.Context, client *ent.Client, params *validators.CreateUserParams) (*ent.User, error) {

	user, err := client.User.
		Create().
		SetEmail(params.Email).
		SetPassword(params.Password).
		SetName(params.Name).
		Save(*ctx)

	return user, err
}
