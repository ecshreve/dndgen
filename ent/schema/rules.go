package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RuleSection holds the schema definition for the RuleSection entity.
type RuleSection struct {
	ent.Schema
}

func (RuleSection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the RuleSection.
func (RuleSection) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc").Annotations(),
	}
}

// Edges of the RuleSection.
func (RuleSection) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the RuleSection.
func (RuleSection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
