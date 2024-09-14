package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MagicSchool holds the schema definition for the MagicSchool entity.
type MagicSchool struct {
	ent.Schema
}

func (MagicSchool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the MagicSchool.
func (MagicSchool) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc"),
	}
}

// Edges of the MagicSchool.
func (MagicSchool) Edges() []ent.Edge {
	return nil
}
