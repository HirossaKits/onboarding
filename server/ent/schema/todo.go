package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Todo struct {
	ent.Schema
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("user_id").NotEmpty(),
		field.Text("title").NotEmpty(),
		field.Text("content").NotEmpty(),
		field.Time("updated_at").Default(time.Now),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("user_todo").Required().Unique(),
	}
}

func (Todo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
