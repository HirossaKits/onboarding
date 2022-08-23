package validators

import (
	"go-chi-api/utils"
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
	params, err := utils.ConvToStruct[CreateUserParams](body)
	return params, err
}

func ValidateUpdateUserParam(body *io.ReadCloser) (UpdateUserParams, error) {
	params, err := utils.ConvToStruct[UpdateUserParams](body)
	return params, err
}
