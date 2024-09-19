package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Alignment holds the schema definition for the Alignment entity.
type Alignment struct {
	ent.Schema
}

func (Alignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Alignment.
func (Alignment) Fields() []ent.Field {
	return []ent.Field{
		field.String("abbr").StructTag(`json:"abbreviation"`),
	}
}
