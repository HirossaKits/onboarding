package service

import (
	"context"
	"go-chi-api/internal/db/ent"
	"go-chi-api/internal/db/ent/user"
	"go-chi-api/internal/validator"

	"github.com/google/uuid"
)

func GetUserById(ctx *context.Context, client *ent.Client, user_id string) (*ent.User, error) {

	uuid, _ := uuid.Parse(user_id)

	user, err := client.User.
		Query().
		Where(user.ID(uuid)).
		Select(user.FieldID, user.FieldName).
		Only(*ctx)

	return user, err
}

func CreateUser(ctx *context.Context, client *ent.Client, params *validator.CreateUserParams) (*ent.User, error) {

	user, err := client.User.
		Create().
		SetEmail(params.Email).
		SetPassword(params.Password).
		SetName(params.Name).
		Save(*ctx)

	return user, err
}

func UpdateUser(ctx *context.Context, client *ent.Client, user_id string, params *validator.UpdateUserParams) (*ent.User, error) {

	uuid, _ := uuid.Parse(user_id)

	user, err := client.User.
		UpdateOneID(uuid).
		SetName(params.Name).
		Save(*ctx)

	return user, err
}

func DeleteUser(ctx *context.Context, client *ent.Client, user_id string) error {

	uuid, _ := uuid.Parse(user_id)

	err := client.User.
		DeleteOneID(uuid).
		Exec(*ctx)

	return err
}
