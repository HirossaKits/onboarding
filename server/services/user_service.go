package services

import (
	"context"
	"go-chi-api/controllers"
	"go-chi-api/ent"
	"go-chi-api/ent/user"
)

func GetUserById(ctx *context.Context, client *ent.Client, param controllers.UserIdParam) (string, error) {

	user, err := client.User.
		Query().
		// Unique().
		Select(user.FieldID, user.FieldName).
		String(*ctx)

	return user, err
}
