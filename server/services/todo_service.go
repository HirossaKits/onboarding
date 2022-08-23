package services

import (
	"context"
	"go-chi-api/ent"
	"go-chi-api/ent/todo"
	"go-chi-api/validators"

	"github.com/google/uuid"
)

func GetTodos(ctx *context.Context, client *ent.Client, params *validators.GetTodosParams) (*ent.Todo, error) {

	uuid, _ := uuid.Parse(params.User_id)

	todo, err := client.Todo.
		Query().
		Where(todo.UserIDEQ(uuid)).
		Select(todo.FieldID,
			todo.FieldTitle,
			todo.FieldContent,
			todo.FieldCreatedAt,
			todo.FieldCreatedAt).
		Only(*ctx)

	return todo, err
}

func GetTodoById(ctx *context.Context, client *ent.Client, todo_id string) ([]*ent.Todo, error) {

	uuid, _ := uuid.Parse(todo_id)

	todo, err := client.Todo.
		Query().
		Where(todo.ID(uuid)).
		Select(todo.FieldID,
			todo.FieldTitle,
			todo.FieldContent,
			todo.FieldCreatedAt,
			todo.FieldCreatedAt).
		All(*ctx)

	return todo, err
}

func CreateTodo(ctx *context.Context, client *ent.Client, params *validators.CreateTodoParams) (*ent.Todo, error) {

	uuid, _ := uuid.Parse(params.User_id)

	todo, err := client.Todo.
		Create().
		SetUserID(uuid).
		SetTitle(params.Title).
		SetContent(params.Content).
		Save(*ctx)

	return todo, err
}

func UpdateTodo(ctx *context.Context, client *ent.Client, todo_id string, params *validators.UpdateTodoParams) (*ent.Todo, error) {

	uuid, _ := uuid.Parse(todo_id)

	todo, err := client.Todo.
		UpdateOneID(uuid).
		SetTitle(params.Title).
		SetContent(params.Content).
		Save(*ctx)

	return todo, err
}

func DeleteTodo(ctx *context.Context, client *ent.Client, user_id string) error {

	uuid, _ := uuid.Parse(user_id)

	err := client.User.
		DeleteOneID(uuid).
		Exec(*ctx)

	return err
}
