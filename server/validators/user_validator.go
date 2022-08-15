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

func ValidateGetUserParam(body *io.ReadCloser) (GetUserParams, error) {
	params, err := ConvToStruct[GetUserParams](body)
	// ここで何かしらのバリデーションを行う
	return params, err
}

func ValidateCreateUserParam(body *io.ReadCloser) (CreateUserParams, error) {
	params, err := ConvToStruct[CreateUserParams](body)
	// ここで何かしらのバリデーションを行う
	return params, err
}
