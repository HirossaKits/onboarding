// Code generated by ent, DO NOT EDIT.

package ent

import (
	"go-chi-api/ent/schema"
	"go-chi-api/ent/todo"
	"go-chi-api/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescUserID is the schema descriptor for user_id field.
	todoDescUserID := todoFields[1].Descriptor()
	// todo.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	todo.UserIDValidator = todoDescUserID.Validators[0].(func(string) error)
	// todoDescTitle is the schema descriptor for title field.
	todoDescTitle := todoFields[2].Descriptor()
	// todo.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	todo.TitleValidator = todoDescTitle.Validators[0].(func(string) error)
	// todoDescContent is the schema descriptor for content field.
	todoDescContent := todoFields[3].Descriptor()
	// todo.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	todo.ContentValidator = todoDescContent.Validators[0].(func(string) error)
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[4].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoFields[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
