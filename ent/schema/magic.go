package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MagicSchool holds the schema definition for the MagicSchool entity.
type MagicSchool struct {
	ent.Schema
}

func (MagicSchool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the MagicSchool.
func (MagicSchool) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc").Annotations(),
	}
}

// Edges of the MagicSchool.
func (MagicSchool) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the MagicSchool.
func (MagicSchool) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
