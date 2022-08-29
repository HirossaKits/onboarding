package validator

import (
	"go-chi-api/pkg/util"
	"io"
)

type GetTodosParams struct {
	User_id string
}

type CreateTodoParams struct {
	User_id string
	Title   string
	Content string
}
type UpdateTodoParams struct {
	Title   string
	Content string
}

func ValidateGetTodosParam(body *io.ReadCloser) (GetTodosParams, error) {
	params, err := util.ConvToStruct[GetTodosParams](body)
	return params, err
}

func ValidateCreateTodoParam(body *io.ReadCloser) (CreateTodoParams, error) {
	params, err := util.ConvToStruct[CreateTodoParams](body)
	return params, err
}

func ValidateUpdateTodoParam(body *io.ReadCloser) (UpdateTodoParams, error) {
	params, err := util.ConvToStruct[UpdateTodoParams](body)
	return params, err
}
