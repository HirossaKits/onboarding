package validators

import (
	"io"
)

type GetTodoParams struct {
	User_id string
}

type CreateTodoParams struct {
	Email    string
	Password string
	Name     string
}
type UpdateTodoParams struct {
	Name string
}

func ValidateCreateTodoParam(body *io.ReadCloser) (CreateUserParams, error) {
	params, err := ConvToStruct[CreateUserParams](body)
	return params, err
}

func ValidateUpdateTodoParam(body *io.ReadCloser) (UpdateUserParams, error) {
	params, err := ConvToStruct[UpdateUserParams](body)
	return params, err
}
