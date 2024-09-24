package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Feature holds the schema definition for the Feature entity.
type Feature struct {
	ent.Schema
}

// Mixin of the Feature.
func (Feature) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Feature.
func (Feature) Fields() []ent.Field {
	return []ent.Field{
		field.Int("level").Positive(),
	}
}
