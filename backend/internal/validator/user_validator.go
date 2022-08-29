package validator

import (
	"go-chi-api/pkg/util"
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
	params, err := util.ConvToStruct[CreateUserParams](body)
	return params, err
}

func ValidateUpdateUserParam(body *io.ReadCloser) (UpdateUserParams, error) {
	params, err := util.ConvToStruct[UpdateUserParams](body)
	return params, err
}
