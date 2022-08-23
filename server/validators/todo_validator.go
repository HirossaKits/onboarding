package validators

import (
	"go-chi-api/utils"
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
	params, err := utils.ConvToStruct[GetTodosParams](body)
	return params, err
}

func ValidateCreateTodoParam(body *io.ReadCloser) (CreateTodoParams, error) {
	params, err := utils.ConvToStruct[CreateTodoParams](body)
	return params, err
}

func ValidateUpdateTodoParam(body *io.ReadCloser) (UpdateTodoParams, error) {
	params, err := utils.ConvToStruct[UpdateTodoParams](body)
	return params, err
}
