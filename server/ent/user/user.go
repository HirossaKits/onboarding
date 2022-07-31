// Code generated by ent, DO NOT EDIT.

package user

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// EdgeUserTodo holds the string denoting the user_todo edge name in mutations.
	EdgeUserTodo = "user_todo"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserTodoTable is the table that holds the user_todo relation/edge.
	UserTodoTable = "todos"
	// UserTodoInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	UserTodoInverseTable = "todos"
	// UserTodoColumn is the table column denoting the user_todo relation/edge.
	UserTodoColumn = "user_user_todo"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPassword,
	FieldName,
	FieldNickname,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)