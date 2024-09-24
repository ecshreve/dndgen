package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tool holds the schema definition for the Tool entity.
type Tool struct {
	ent.Schema
}

// Fields of the Tool.
func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.String("tool_category"),
		field.Strings("desc").Optional(),
	}
}

// Edges of the Tool.
func (Tool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("tool").
			Unique().
			Required(),
	}
}
