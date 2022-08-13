package validators

import "io"

type userValidator struct{}

type GetUserParam struct {
	user_id string
}

func NewUserValidator() *userValidator {
	return &userValidator{}
}

func (*userValidator) ValidateGetUserParam(body []byte) {

}

func Validatep[T any](*io.ReadCloser) {

}
