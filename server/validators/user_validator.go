package validators

import (
	"fmt"
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
	// TODO:ここで何かしらのバリデーションを行う
	return params, err
}

func ValidateCreateUserParam(body *io.ReadCloser) (CreateUserParams, error) {

	fmt.Println(*body)
	params, err := ConvToStruct[CreateUserParams](body)
	fmt.Println("Hey!")
	fmt.Println(params)
	// TODO:ここで何かしらのバリデーションを行う
	return params, err
}
