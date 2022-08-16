package validators

import (
	"io"
)

type GetUserParams struct {
	User_id string
}

type CreateUserParams struct {
	Email    string
	Password string
	Name     string
}
type UpdateUserParams struct {
	Name string
}

func ValidateCreateUserParam(body *io.ReadCloser) (CreateUserParams, error) {
	params, err := ConvToStruct[CreateUserParams](body)
	return params, err
}

func ValidateUpdateUserParam(body *io.ReadCloser) (UpdateUserParams, error) {
	params, err := ConvToStruct[UpdateUserParams](body)
	return params, err
}
