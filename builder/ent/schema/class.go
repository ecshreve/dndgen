package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

func (Class) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return nil
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("characters", Character.Type),
	}
}
