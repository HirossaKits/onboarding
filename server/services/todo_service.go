package services

import (
	"context"
	"go-chi-api/ent"
	"go-chi-api/ent/user"
	"go-chi-api/validators"

	"github.com/google/uuid"
)

func GetTodoById(ctx *context.Context, client *ent.Client, user_id string) (*ent.User, error) {

	uuid, _ := uuid.Parse(user_id)

	user, err := client.User.
		Query().
		Where(user.ID(uuid)).
		Select(user.FieldID, user.FieldName).
		Only(*ctx)

	return user, err
}

func CreateTodo(ctx *context.Context, client *ent.Client, params *validators.CreateUserParams) (*ent.User, error) {

	user, err := client.User.
		Create().
		SetEmail(params.Email).
		SetPassword(params.Password).
		SetName(params.Name).
		Save(*ctx)

	return user, err
}

func UpdateTodo(ctx *context.Context, client *ent.Client, user_id string, params *validators.UpdateUserParams) (*ent.User, error) {

	uuid, _ := uuid.Parse(user_id)

	user, err := client.User.
		UpdateOneID(uuid).
		SetName(params.Name).
		Save(*ctx)

	return user, err
}

func DeleteTodo(ctx *context.Context, client *ent.Client, user_id string) error {

	uuid, _ := uuid.Parse(user_id)

	err := client.User.
		DeleteOneID(uuid).
		Exec(*ctx)

	return err
}
